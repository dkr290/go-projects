package chartio

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

type ChartConfig struct {
	ChartName        string
	ChartVersion     string
	ChartDir         string
	LocalRegisrtry   string
	RegistryUser     string
	RegistryPassword string
}
type ChartMetadata struct {
	AppVersion string `yaml:"appVersion"`
}

func New(
	chartName, chartVersion, chartlocalDir, localRegistry, user, password string,
) *ChartConfig {
	return &ChartConfig{
		ChartName:        chartName,
		ChartVersion:     chartVersion,
		ChartDir:         chartlocalDir,
		LocalRegisrtry:   localRegistry,
		RegistryUser:     user,
		RegistryPassword: password,
	}
}

func (c *ChartConfig) PullChart() (string, error) {
	cmd := exec.Command(
		"helm",
		"pull",
		c.ChartName,
		"--version",
		c.ChartVersion,
		"--destination",
		c.ChartDir,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	// Helm saves the chart as "chartname-version.tgz", so we determine the filename
	tgzFile := fmt.Sprintf(
		"%s/%s-%s.tgz",
		c.ChartDir,
		strings.Split(c.ChartName, "/")[1],
		c.ChartVersion,
	)
	return tgzFile, nil
}

// Get appVersion from Chart.yaml
func (c *ChartConfig) GetAppVersion(chartPath string) (string, error) {
	data, err := os.ReadFile(chartPath)
	if err != nil {
		return "", fmt.Errorf("failed to read Chart.yaml: %v", err)
	}

	var chartMeta ChartMetadata
	if err := yaml.Unmarshal(data, &chartMeta); err != nil {
		return "", fmt.Errorf("failed to parse Chart.yaml: %v", err)
	}

	if chartMeta.AppVersion == "" {
		return "", fmt.Errorf("appVersion is empty in Chart.yaml")
	}

	return chartMeta.AppVersion, nil
}

// Extract images from values.yaml with tag resolution
func (c *ChartConfig) ExtractImages(valuesPath, appVersion string) ([]string, error) {
	data, err := os.ReadFile(valuesPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read values.yaml: %v", err)
	}

	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return nil, fmt.Errorf("failed to parse values.yaml: %v", err)
	}

	var images []string
	traverseNodes(&root, appVersion, &images)

	if len(images) == 0 {
		return nil, fmt.Errorf("no images found in values.yaml")
	}

	return images, nil
}

// Recursive YAML node traversal
func traverseNodes(node *yaml.Node, appVersion string, images *[]string) {
	switch node.Kind {
	case yaml.DocumentNode:
		for _, n := range node.Content {
			traverseNodes(n, appVersion, images)
		}
	case yaml.MappingNode:
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			value := node.Content[i+1]

			if key.Value == "repository" {
				repo := value.Value
				tag := appVersion

				// Search for tag in the same mapping node
				for j := 0; j < len(node.Content); j += 2 {
					if node.Content[j].Value == "tag" {
						tagNode := node.Content[j+1]
						if tagNode.Value != "" && tagNode.Value != "null" {
							tag = tagNode.Value
						}
						break
					}
				}

				*images = append(*images, fmt.Sprintf("%s:%s", repo, tag))
			} else {
				traverseNodes(value, appVersion, images)
			}
		}
	case yaml.SequenceNode:
		for _, n := range node.Content {
			traverseNodes(n, appVersion, images)
		}
	case yaml.AliasNode:
		traverseNodes(node.Alias, appVersion, images)
	}
}
