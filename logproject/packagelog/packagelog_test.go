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

// type testWriter struct {
// 	contents string
// }

// func (t *testWriter) Write(p []byte)(n int,err error){
// 	t.contents = t.contents+string(p)
// 	return len(p),nil
// }

// const (
// 	debugMessage = "This is from debug level"
// 	infoMessage  = "This is from the info level"
// 	errorMessage = "This is from the error message"
// )

// func TestLogger_DebugInfoError(t *testing.T) {

// 	type testCase struct {
// 		level    packagelog.Level
// 		expected string
// 	}

// 	tt := map[string]testCase{
// 		"debug": {
// 			level:    packagelog.LevelDebug,
// 			expected: debugMessage + "\n" + infoMessage + "\n" + errorMessage + "\n",
// 		},
// 		"info": {
// 			level:    packagelog.LevelInfo,
// 			expected: infoMessage + "\n" + errorMessage + "\n",
// 		},
// 		"error": {
// 			level:    packagelog.LevelError,
// 			expected: errorMessage + "\n",
// 		},
// 	}

// 	for name,tc := range tt{
// 		t.Run(name func (t *&testing.T{
// 			tw := &testWriter{}
// 			 testedLogger := packagelog.New(tc.level,packagelog.)
// 		}){

// 		}

// }
