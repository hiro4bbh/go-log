package golog

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/hiro4bbh/go-assert"
)

func TestParseLogLevel(t *testing.T) {
	goassert.New(t, DEBUG).EqualWithoutError(ParseLogLevel("DEBUG"))
	goassert.New(t, DEBUG).EqualWithoutError(ParseLogLevel("debug"))
	goassert.New(t, INFO).EqualWithoutError(ParseLogLevel("INFO"))
	goassert.New(t, WARN).EqualWithoutError(ParseLogLevel("WARN"))
	goassert.New(t, ERROR).EqualWithoutError(ParseLogLevel("ERROR"))
	goassert.New(t, "illegal LogLevel: \"illegal\"").ExpectError(ParseLogLevel("illegal"))
	goassert.New(t, "illegal LogLevel: \"\"").ExpectError(ParseLogLevel(""))
}

func TestLogLevelString(t *testing.T) {
	goassert.New(t, "LogLevel(0)").Equal(illegalLogLevel.String())
	goassert.New(t, "DEBUG").Equal(DEBUG.String())
	goassert.New(t, "INFO").Equal(INFO.String())
	goassert.New(t, "WARN").Equal(WARN.String())
	goassert.New(t, "ERROR").Equal(ERROR.String())
}

func TestLoggerDefault(t *testing.T) {
	exitcode := 0
	ExitHandler = func(code int) {
		exitcode = code
	}
	var buf bytes.Buffer
	logger := New(&buf, &Parameters{})
	goassert.New(t, DefaultMinLevel).Equal(logger.MinLevel())
	goassert.New(t, DefaultName).Equal(logger.Name())
	goassert.New(t, DefaultTimeFormat).Equal(logger.TimeFormat())
	goassert.New(t, false).Equal(logger.Color())
	logger.Debugf("Hello from DEBUG level")
	logger.Infof("Hello from INFO level")
	logger.Warnf("Hello from WARN level")
	logger.Errorf("Hello from ERROR level")
	logger.Writer().Write([]byte("hello\n"))
	goassert.New(t, true).Equal(regexp.MustCompile("^ WARN.*Hello from WARN level\nERROR.*Hello from ERROR level\nhello\n$").MatchString(buf.String()))
	logger.SetMinLevel(DEBUG)
	logger.Debugf("Hello from DEBUG level")
	logger.Infof("Hello from INFO level")
	goassert.New(t, true).Equal(regexp.MustCompile("^ WARN.*Hello from WARN level\nERROR.*Hello from ERROR level\nhello\nDEBUG.*Hello from DEBUG level\n INFO.*Hello from INFO level\n$").MatchString(buf.String()))
	goassert.New(t, 0).Equal(exitcode)
}

func TestLoggerFatalf(t *testing.T) {
	exitcode := 0
	ExitHandler = func(code int) {
		exitcode = code
	}
	var buf bytes.Buffer
	logger := New(&buf, &Parameters{})
	logger.Fatalf("Hello from ERROR level")
	goassert.New(t, true).Equal(regexp.MustCompile("^ERROR.*Hello from ERROR level\nERROR.*hit FATAL error: exiting with status code 1\n$").MatchString(buf.String()))
	goassert.New(t, 1).Equal(exitcode)
}

func TestLoggerColorModeAlways(t *testing.T) {
	exitcode := 0
	ExitHandler = func(code int) {
		exitcode = code
	}
	var buf bytes.Buffer
	logger := New(&buf, &Parameters{Name: "always", MinLevel: DEBUG, ColorMode: "always"})
	goassert.New(t, true).Equal(logger.Color())
	logger.Debugf("Hello from DEBUG level")
	logger.Infof("Hello from INFO level")
	logger.Warnf("Hello from WARN level")
	logger.Errorf("Hello from ERROR level")
	goassert.New(t, true).Equal(regexp.MustCompile("^\033\\[1mDEBUG\033\\[0m.*Hello from DEBUG level\n\033\\[36m INFO\033\\[0m.*Hello from INFO level\n\033\\[33m WARN\033\\[0m.*Hello from WARN level\n\033\\[31mERROR\033\\[0m.*Hello from ERROR level\n$").MatchString(buf.String()))
	goassert.New(t, 0).Equal(exitcode)
}

func TestLoggerColorModeNever(t *testing.T) {
	exitcode := 0
	ExitHandler = func(code int) {
		exitcode = code
	}
	var buf bytes.Buffer
	logger := New(&buf, &Parameters{Name: "never", MinLevel: DEBUG, ColorMode: "never"})
	goassert.New(t, false).Equal(logger.Color())
	logger.Debugf("Hello from DEBUG level")
	logger.Infof("Hello from INFO level")
	logger.Warnf("Hello from WARN level")
	logger.Errorf("Hello from ERROR level")
	goassert.New(t, true).Equal(regexp.MustCompile("^DEBUG.*Hello from DEBUG level\n INFO.*Hello from INFO level\n WARN.*Hello from WARN level\nERROR.*Hello from ERROR level\n$").MatchString(buf.String()))
	goassert.New(t, 0).Equal(exitcode)
}
