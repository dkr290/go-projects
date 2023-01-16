package main

import (
	"os"

	"github.com/dkr290/go-projects/logproject/packagelog"
)

func main() {
	lgr := packagelog.New(packagelog.LevelInfo, os.Stdout)

	lgr.Debugf("This is from debug level %d", 0)
	lgr.Infof("Some message concerning info %s", "level")
	lgr.Errorf("Errors are values. Documentations is for %s.", "users")
}
