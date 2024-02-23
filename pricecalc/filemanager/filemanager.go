package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

// just reading so no need a pointer
func (f FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(f.InputFilePath)

	if err != nil {
		file.Close()
		return nil, errors.New("an error opening the file occured" + err.Error())
	}
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("reading the file content failed" + err.Error())
	}

	file.Close()
	return lines, nil
}

func (f FileManager) WriteResult(data any) error {

	file, err := os.Create(f.OutputFilePath)
	if err != nil {
		return errors.New("failed to create a file")
	}
	if err := json.NewEncoder(file).Encode(data); err != nil {
		file.Close()
		return errors.New("failed to conver data to json")
	}
	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {

	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
