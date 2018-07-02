package golog

import (
	"fmt"
	"io"

	"github.com/hiro4bbh/go-term"
)

// Style has the font color and decoration information.
type Style uint64

// The font decoration styles.
const (
	Normal    = Style(0x00)
	Bold      = Style(1 << 0)
	Underline = Style(1 << 1)
)

// The foreground basic colors supported by ANSI escape codes.
const (
	FgBlack   = Style(0 << 8)
	FgRed     = Style(1 << 8)
	FgGreen   = Style(2 << 8)
	FgYellow  = Style(3 << 8)
	FgBlue    = Style(4 << 8)
	FgMagenta = Style(5 << 8)
	FgCyan    = Style(6 << 8)
	FgGray    = Style(7 << 8)
)

// Bold returns the bolded Style if bold is true, otherwise un-bolded Style.
func (style Style) Bold(bold bool) Style {
	if bold {
		return style | Bold
	}
	return style & ^Bold
}

// IsBold returns true for bold Styles.
func (style Style) IsBold() bool {
	return (style & Bold) != 0
}

// Underline returns the underlined Style if bold is true, otherwise un-underlined Style.
func (style Style) Underline(underline bool) Style {
	if underline {
		return style | Underline
	}
	return style & ^Underline
}

// IsUnderline returns true for underline Styles.
func (style Style) IsUnderline() bool {
	return (style & Underline) != 0
}

// FgColor returns the foreground color information.
func (style Style) FgColor() Style {
	return style & 0x700
}

// Fprintf prints the result of Sprintf to the given writer.
// If the writer is not terminal, any style will be ignored.
func (style Style) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if goterm.IsTerminal(w) {
		return fmt.Fprint(w, style.Sprintf(format, a...))
	}
	return fmt.Fprint(w, fmt.Sprintf(format, a...))
}

// SetFgColor returns the given Style with the specified color.
// If the specified color is illegal, then nothing are affected.
func (style Style) SetFgColor(color Style) Style {
	return (style & ^Style(0x700)) | (color & Style(0x700))
}

// Sprintf formats according to a format specifier and returns the resulting string with the specified style.
func (style Style) Sprintf(format string, a ...interface{}) string {
	start, end := "", ""
	if style.IsBold() {
		start += "\033[1m"
	}
	if style.IsUnderline() {
		start += "\033[4m"
	}
	switch style.FgColor() {
	case FgRed:
		start += "\033[31m"
	case FgGreen:
		start += "\033[32m"
	case FgYellow:
		start += "\033[33m"
	case FgBlue:
		start += "\033[34m"
	case FgMagenta:
		start += "\033[35m"
	case FgCyan:
		start += "\033[36m"
	case FgGray:
		start += "\033[37m"
	}
	if start != "" {
		end = "\033[0m"
	}
	return fmt.Sprintf(start+format+end, a...)
}
