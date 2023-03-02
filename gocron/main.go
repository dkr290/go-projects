package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jasonlvhit/gocron"
)

func task() {

	filename := GetFilenameDate()

	cmd := exec.Command("logcli", "query", "--since=1h", "--limit=500000000000000", `'{namespace="loki"}'`)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go io.Copy(writer, stdoutPipe)
	cmd.Wait()

}

func main() {
	s := gocron.NewScheduler()
	s.Every(1).Hours().Do(task)

	<-s.Start()
}

func GetFilenameDate() string {
	// Use layout string for time format.
	const layout = "01-02-2006"
	// Place now in the string.
	t := time.Now()
	return "log-" + t.Format(layout) + ".txt"
}
