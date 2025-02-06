package main

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
)

var ErrorFileNotFount = errors.New("The working file was not found")

func main() {
	err := createBackup("note.txt", "note.txt.backup")
	if err != nil {
		log.Println(err)
	}
}

func buffReadFile(file string, buffsize int) (bytes.Buffer, error) {
	var nbuff bytes.Buffer
	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return bytes.Buffer{}, ErrorFileNotFount
		}
		return bytes.Buffer{}, err
	}
	buff := make([]byte, buffsize)
	for {
		n, err := f.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}
		if n > 0 {
			nbuff.WriteString(string(buff[:n]))
		}
	}

	return nbuff, nil
}

func createBackup(working, backup string) error {
	// check if the working file exists

	_, err := os.Stat(working)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrorFileNotFount
		}
		return err
	}
	stringbuff, err := buffReadFile(working, 50)
	if err != nil {
		return err
	}
	err = os.WriteFile(backup, stringbuff.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}
