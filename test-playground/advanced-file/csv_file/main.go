package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	ch := make(chan []byte)
	go func() {
		err := readCsv("file.csv", ch, 4)
		if err != nil {
			log.Println(err)
		}
	}()

	for _, v := range <-ch {
		fmt.Print(string(v))
	}
}

func readCsv(filename string, ch chan []byte, size int) error {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("The file does not exists " + err.Error())
		}
		return err
	}

	buff := make([]byte, size)
	var allBuff bytes.Buffer
	for {
		n, err := f.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			allBuff.WriteString(string(buff[:n]))
		}
	}
	ch <- allBuff.Bytes()

	return nil
}
