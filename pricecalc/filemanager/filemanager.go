package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

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
