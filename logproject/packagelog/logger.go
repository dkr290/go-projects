package packagelog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used
type Logger struct {
	treshold Level
	output   io.Writer
}

// Debugf formats and prints a message of the log levels of debug or lower
func (l *Logger) Debugf(format string, args ...any) {
	//making sure we safetly write to the output
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.treshold <= LevelDebug {
		_, _ = fmt.Fprintf(l.output, format, args...)
	}
}

// Infof formats and prints messages of log level of info or lower

func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.treshold <= LevelInfo {
		_, _ = fmt.Printf(format+"\n", args...)

	}

}

func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.treshold <= LevelError {
		_, _ = fmt.Printf(format+"\n", args...)
	}

}

// New returns you a logger , ready to log
// The default output is the stdout
func New(treshold Level, output io.Writer) *Logger {
	return &Logger{
		treshold: treshold,
		output:   output,
	}
}
