package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

var Directory string
var DaysToremove int
var fileExt string
var help = flag.Bool("help", false, "Show help")

func main() {
	flag.StringVar(&Directory, "dir", "", "The Directory for search to delete old files")
	flag.IntVar(&DaysToremove, "days", 0, "The days value for files to keep - more then x days difference delete")
	flag.StringVar(&fileExt, "ext", "", "The files extension to remove like ex .log")

	// Parse the flag
	flag.Parse()

	// Usage Demo
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if Directory == "" || DaysToremove == 0 || fileExt == "" {
		flag.Usage()
		os.Exit(0)
	}

	for _, s := range find(Directory, fileExt) {
		removeOldLogs(DaysToremove, s)

	}
}

func find(dirpath, ext string) []string {
	var a []string
	filepath.WalkDir(dirpath, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func removeOldLogs(olddays int, f string) {

	timeNow := time.Now()

	stat, err := os.Stat(f)
	if err != nil {
		fmt.Println("Cannot determine the file")
	}
	fModTime := stat.ModTime()
	loc, _ := time.LoadLocation("UTC")
	l1 := fModTime.In(loc)
	l2 := timeNow.In(loc)
	diff := l2.Sub(l1)

	if diff.Hours() > float64(olddays*24) {
		fmt.Println("Deleting file", f, "With difftime in hours", diff.Hours())
		e := os.Remove(f)
		if e != nil {
			log.Printf("Cannot remove file %s %v", f, e)
		}
	}

}
