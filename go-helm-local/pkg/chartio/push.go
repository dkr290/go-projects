package chartio

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-helm-local/pkg/helpers"
	"io"
	"log/slog"
	"strings"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

// Struct to hold the parsed JSON data
type DockerProgress struct {
	Status         string `json:"status"`
	Progress       string `json:"progress"`
	ID             string `json:"id"`
	ProgressDetail struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail"`
}

func (c *ChartConfig) PushImageToRegistry(
	cli *client.Client,
	dockerImage string,
	localRegistry string,
) (string, error, bool) {
	// Extract image and tag
	imageNameParts := strings.Split(dockerImage, ":")
	if imageNameParts[0] == "" || imageNameParts[0] == "null" {
		return "", nil, true
	}
	// Replace remote registry with local registry
	newImage := strings.Replace(dockerImage, strings.Split(dockerImage, "/")[0], localRegistry, 1)
	// // Pull the dockerImage
	helpers.Logging("Pulling docker image "+dockerImage, "info", nil)
	pullResponse, err := cli.ImagePull(context.Background(), dockerImage, image.PullOptions{})
	if err != nil {
		return "", err, false
	}
	defer pullResponse.Close()
	// Use formatted output instead of raw io.Copy
	if err := formatDockerProgress(pullResponse); err != nil {
		slog.Error("error from formatter docker output", "error", err)
	}
	// Verify image pull before tagging
	_, _, err = cli.ImageInspectWithRaw(context.Background(), dockerImage)
	if err != nil {
		return "", fmt.Errorf("image %s not found after pull: %w", dockerImage, err), false
	}
	// Tag the image using Docker SDK
	err = cli.ImageTag(context.Background(), dockerImage, newImage)
	if err != nil {
		return "", fmt.Errorf(
			"failed to tag image %s as %s: %w",
			dockerImage,
			newImage,
			err,
		), false
	}
	// Create the authentication configuration
	authConfig := registry.AuthConfig{
		Username: c.RegistryUser,
		Password: c.RegistryPassword,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return "", err, false
	}
	authStr := base64.StdEncoding.EncodeToString(encodedJSON)
	// Push to the local registry
	// Verify tag before pushing
	_, _, err = cli.ImageInspectWithRaw(context.Background(), newImage)
	if err != nil {
		return "", fmt.Errorf(
			"failed to inspect image %s before pushing: %w",
			newImage,
			err,
		), false
	}
	fmt.Println("Pushing image:", newImage)
	helpers.Logging("Pushing image "+newImage, "info", nil)
	pushResponse, err := cli.ImagePush(context.Background(), newImage, image.PushOptions{
		RegistryAuth: authStr,
	})
	if err != nil {
		return "", err, false
	}
	defer pushResponse.Close()
	// Use formatted output instead of raw io.Copy
	if err := formatDockerProgress(pushResponse); err != nil {
		slog.Error("error from formatter docker output", "error", err)
	}

	fmt.Println("Deleting local image")
	helpers.Logging("Deleting local image "+newImage, "info", nil)
	_, err = cli.ImageRemove(context.Background(), newImage, image.RemoveOptions{})
	if err != nil {
		fmt.Println("error removing image ", newImage)
	}
	_, err = cli.ImageRemove(context.Background(), dockerImage, image.RemoveOptions{})
	if err != nil {
		fmt.Println("error removing image ", newImage)
	}

	return newImage, nil, false
}

func formatDockerProgress(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		var progress DockerProgress

		// Try to parse the JSON
		if err := json.Unmarshal([]byte(line), &progress); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		// Display formatted output
		if progress.ID != "" {
			fmt.Printf("[%-12s] %-25s %s\n", progress.ID, progress.Status, progress.Progress)
		} else {
			fmt.Printf("[%-12s] %-25s\n", "-", progress.Status)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading input %v", err)
	}

	return nil
}
