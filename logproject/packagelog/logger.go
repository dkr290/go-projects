package packagelog

import (
	"fmt"
	"io"
	"os"
	"time"
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

		l.logf(LevelDebug, format, args...)
	}
}

// Infof formats and prints messages of log level of info or lower

func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.treshold <= LevelInfo {
		l.logf(LevelInfo, format, args...)

	}

}

func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.treshold <= LevelError {
		l.logf(LevelError, format, args...)
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

// logf prints the message to the output
func (l *Logger) logf(lvl Level, format string, args ...any) {

	if l.treshold > lvl {
		return
	}

	var NewLevel string
	switch lvl {
	case 0:
		NewLevel = "Debug"

	case 1:
		NewLevel = "Info"
	case 2:
		NewLevel = "Error"
	}

	t := time.Now()
	s := t.Format(time.RFC850)
	_, _ = fmt.Fprintf(l.output, "time: %s\n", s)
	_, _ = fmt.Fprintf(l.output, "level: %v\n", NewLevel)
	_, _ = fmt.Fprintf(l.output, "message: "+format+"\n", args...)
	fmt.Println("")
}
