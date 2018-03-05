package golog

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// LogLevel is the level used in each logged line.
type LogLevel int

const (
	illegalLogLevel = LogLevel(0)
)
const (
	DEBUG = LogLevel(1)
	INFO  = LogLevel(2)
	WARN  = LogLevel(3)
	ERROR = LogLevel(4)
)

// ParseLogLevel parses LogLevel from the given str.
//
// This function returns error in parsing.
func ParseLogLevel(str string) (LogLevel, error) {
	switch strings.ToUpper(str) {
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	}
	return illegalLogLevel, fmt.Errorf("illegal LogLevel: %q", str)
}

// String returns the string representation of LogLevel.
func (level LogLevel) String() string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	}
	return fmt.Sprintf("LogLevel(%d)", level)
}

// Parameters has the parameter for Logger.
type Parameters struct {
	// Name is logging title.
	Name string
	// MinLevel is the minimum logged priority.
	MinLevel LogLevel
	// TimeFormat is used in logging timestamp.
	// This format is based on time.Time.Format.
	TimeFormat string
	// ColorMode specifies the coloring.
	// If this is "always", then the coloring will be enabled always.
	// If this is "never", then the coloring will be disabled always.
	// Otherwise, the coloring will be enabled if IsTerminal(out) is true.
	ColorMode string
}

// DefaultName is the default value of Parameters.Name.
var DefaultName = ""

// DefaultMinLevel is the default value of Parameters.MinLevel.
// This is WARN, but, if the environment variable GOLOG_MINLEVEL is defined and legal, its value is used.
// Do not assign illegal LogLevel, because there is no check for this.
var DefaultMinLevel = WARN

// DefaultTimeFormat is the default value of Parameters.TimeFormat.
var DefaultTimeFormat = "2006/01/02T15:04:05"

// ExitHandler is the exit handler used in FATAL logging.
var ExitHandler = os.Exit

// Logger is the logging manager.
type Logger struct {
	out    io.Writer
	params *Parameters
}

// New returns a new Logger with the specified params.
// If params is nil, then it is filled with default values.
// If params has some zero-value fields, then those are replaced with the corresponding default values.
func New(out io.Writer, params *Parameters) *Logger {
	if params == nil {
		params = &Parameters{
			Name:       DefaultName,
			MinLevel:   DefaultMinLevel,
			TimeFormat: DefaultTimeFormat,
		}
	}
	if params.Name == "" {
		params.Name = DefaultName
	}
	if params.MinLevel == illegalLogLevel {
		params.MinLevel = DefaultMinLevel
	}
	if params.TimeFormat == "" {
		params.TimeFormat = DefaultTimeFormat
	}
	return &Logger{
		out:    out,
		params: params,
	}
}

// Color returns true if the coloring is enabled.
func (logger *Logger) Color() bool {
	switch logger.params.ColorMode {
	case "always":
		return true
	case "never":
		return false
	}
	return IsTerminal(logger.out)
}

// MinLevel returns the minimum logging level.
func (logger *Logger) MinLevel() LogLevel {
	return logger.params.MinLevel
}

// SetMinLevel sets the minimum logging level.
func (logger *Logger) SetMinLevel(level LogLevel) {
	logger.params.MinLevel = level
}

// Name returns the name.
func (logger *Logger) Name() string {
	return logger.params.Name
}

// TimeFormat returns the time format.
func (logger *Logger) TimeFormat() string {
	return logger.params.TimeFormat
}

// Writer returns the underlying writer.
func (logger *Logger) Writer() io.Writer {
	return logger.out
}

// Debugf writes the formatted message with DEBUG level.
func (logger *Logger) Debugf(format string, a ...interface{}) {
	logger.Log(DEBUG, fmt.Sprintf(format, a...))
}

// Infof writes the formatted message with INFO level.
func (logger *Logger) Infof(format string, a ...interface{}) {
	logger.Log(INFO, fmt.Sprintf(format, a...))
}

// Warnf writes the formatted message with WARN level.
func (logger *Logger) Warnf(format string, a ...interface{}) {
	logger.Log(WARN, fmt.Sprintf(format, a...))
}

// Errorf writes the formatted message with ERROR level.
func (logger *Logger) Errorf(format string, a ...interface{}) {
	logger.Log(ERROR, fmt.Sprintf(format, a...))
}

// Fatalf is Errorf followed by ExitHandler(1).
func (logger *Logger) Fatalf(format string, a ...interface{}) {
	logger.Errorf(format, a...)
	logger.Errorf("hit FATAL error: exiting with status code 1")
	ExitHandler(1)
}

// Log writes the given string with the given logging level.
func (logger *Logger) Log(level LogLevel, msg string) {
	if level < logger.params.MinLevel {
		return
	}
	style := Normal
	if logger.Color() {
		switch level {
		case DEBUG:
			style = style.Bold(true)
		case INFO:
			style = style.SetFgColor(FgCyan)
		case WARN:
			style = style.SetFgColor(FgYellow)
		case ERROR:
			style = style.SetFgColor(FgRed)
		}
	}
	timestamp := time.Now().Format(logger.params.TimeFormat)
	fmt.Fprintf(logger.out, "%s %s %s%s\n", style.Sprintf("%5s", level.String()), timestamp, logger.params.Name, strings.TrimSpace(msg))
}

// Null is the logger writing to ioutil.Discard.
var Null = New(ioutil.Discard, nil)

func init() {
	level, err := ParseLogLevel(os.Getenv("GOLOG_MINLEVEL"))
	if err == nil {
		DefaultMinLevel = level
	}
}
