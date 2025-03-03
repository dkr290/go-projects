package helpers

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// extractTGZ extracts a .tgz Helm chart into the destination directory
func ExtractTGZ(tgzFile, destDir string) (string, error) {
	// Open the Helm .tgz file
	file, err := os.Open(tgzFile)
	if err != nil {
		return "", fmt.Errorf("failed to open tgz file: %w", err)
	}
	defer file.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	// Create a tar reader
	tarReader := tar.NewReader(gzipReader)

	var extractedFolder string

	// Extract files from the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read tar archive: %w", err)
		}

		targetPath := filepath.Join(destDir, header.Name)

		// Capture the top-level directory name
		if extractedFolder == "" {
			extractedFolder = filepath.Join(destDir, strings.Split(header.Name, "/")[0])
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(targetPath, 0755); err != nil {
				return "", fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg:
			// Ensure parent directory exists
			if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
				return "", fmt.Errorf("failed to create parent directory: %w", err)
			}

			outFile, err := os.Create(targetPath)
			if err != nil {
				return "", fmt.Errorf("failed to create file: %w", err)
			}
			_, err = io.Copy(outFile, tarReader)
			outFile.Close()
			if err != nil {
				return "", fmt.Errorf("failed to write file: %w", err)
			}
		}
	}

	// Ensure we found a valid extracted folder
	if extractedFolder == "" {
		return "", fmt.Errorf("failed to determine extracted folder")
	}

	fmt.Println("Extracted Helm chart to:", extractedFolder)
	return extractedFolder, nil
}

func Logging(message string, t string, err error) {
	var logger zerolog.Logger
	if t == "error" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		logger.Error().Err(err).Msg(message)
	}

	if t == "info" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Info().Msg(message)
	}
}
