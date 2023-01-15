package packagelog_test

import (
	"os"
	"testing"

	"github.com/dkr290/go-projects/logproject/packagelog"
)

func TestPackageLog_Debugf(t *testing.T) {
	debugLogger := packagelog.New(packagelog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello %s", "world")
}
