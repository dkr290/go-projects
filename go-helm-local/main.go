package main

import (
	"flag"
	"fmt"
	"go-helm-local/pkg/chartio"
	"go-helm-local/pkg/helpers"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/client"
)

var (
	chartName        *string
	version          *string
	destDir          *string
	localRegistry    *string
	registryUsername *string
	registryPassword *string
)

func main() {
	// Define flags for the Helm chart and other parameters
	chartName = flag.String(
		"chart",
		"",
		"Name of the Helm chart to pull (e.g., grafana/loki)",
	)
	version = flag.String("version", "", "Version of the Helm chart to pull (e.g., 6.27.0)")
	destDir = flag.String("dest", "./temp", "Directory to download the Helm chart to")
	localRegistry = flag.String(
		"registry",
		"",
		"Local container registry to push the images to",
	)
	registryUsername = flag.String(
		"username",
		"",
		"Local container registry username",
	)
	registryPassword = flag.String(
		"password",
		"",
		"Local container registry password",
	)

	// Parse the command-line flags
	CmdLineparams()

	c := chartio.New(
		*chartName,
		*version,
		*destDir,
		*localRegistry,
		*registryUsername,
		*registryPassword,
	)
	// Download the Helm chart
	tgzfile, err := c.PullChart()
	if err != nil {
		helpers.Logging("Error pulling helm chart", "error", err)
		os.Exit(1)
	}
	// Initialize Docker client
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal(err)
	}

	// Step 2: Extract Helm Chart
	extractedDir, err := helpers.ExtractTGZ(tgzfile, *destDir)
	if err != nil {
		helpers.Logging("Error extracting Helm Chart", "error", err)
		os.Exit(1)
	}

	// Step 3: Locate Chart.yaml and values.yaml
	chartPath := filepath.Join(extractedDir, "Chart.yaml")
	valuesPath := filepath.Join(extractedDir, "values.yaml")
	// Parse the values.yaml for image references
	appVersion, err := c.GetAppVersion(chartPath)
	if err != nil {
		log.Fatalln("Error reading Chart.yaml", err)
	}
	images, err := c.ExtractImages(valuesPath, appVersion)
	if err != nil {
		log.Fatalln("Error extracting images: ", err)
	}

	var pushedimageSummary []string
	var skipped bool
	helpers.Logging("Resolved images:", "info", nil)
	for _, img := range images {
		// Push the image to the local registry
		img, err, skipped = c.PushImageToRegistry(cli, img, *localRegistry)
		if !skipped {
			if err != nil {
				m := fmt.Sprintf("Failed to push image %s: \n", img)
				helpers.Logging(m, "error", err)
			} else {
				m := fmt.Sprintf("Successfully pushed %s to %s\n", img, *localRegistry)
				helpers.Logging(m, "info", nil)
				pushedimageSummary = append(pushedimageSummary, img)
			}
		}
	}

	fmt.Println("The pushed image summary")
	fmt.Println("========================")
	for _, smImage := range pushedimageSummary {
		fmt.Printf("Image name: %s\n", smImage)
	}
}

func CmdLineparams() {
	// Define command-line flags with default values
	help := flag.Bool("help", false, "Show usage information")
	// Parse command-line flags
	flag.Parse()
	// Display help if the flag is set
	if *help {
		showUsage()
	}

	// Ensure required parameters are provided
	if *chartName == "" {
		fmt.Printf("-chart not supplied %s.\n", *chartName)
		showUsage()
	}

	if *version == "" {
		fmt.Printf("-version not supplied  %s. \n", *version)
		showUsage()
	}
	if *destDir == "" {
		fmt.Printf("-dest not supplied  %s. \n", *destDir)
		showUsage()
	}
	if *localRegistry == "" {
		fmt.Printf("-registry not supplied %s. \n", *localRegistry)
		showUsage()
	}
	if *registryUsername == "" {
		fmt.Printf("-username not supplied %s. \n", *registryUsername)
		showUsage()
	}

	if *registryPassword == "" {
		fmt.Printf("-password not supplied %s. \n", *registryPassword)
		showUsage()
	}
}

// showUsage prints usage information and exits
func showUsage() {
	fmt.Println("Usage:")
	fmt.Println("  -chart string the chart like grafana/loki")
	fmt.Println("  -version  string version like 6.27.0 -registry localhost:5000")
	fmt.Println("  -dest string like -dest ./temp")
	fmt.Println("  -registry string like containerregistry.azurecr.io")
	fmt.Println("  -username the registry username")
	fmt.Println("  -password the registry password")
	fmt.Println("  -help Show this help message")

	os.Exit(0)
}
