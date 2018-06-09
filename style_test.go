package golog

import (
	"bytes"
	"testing"

	"github.com/hiro4bbh/go-assert"
)

func TestFgColor(t *testing.T) {
	goassert.New(t, FgRed).Equal(FgRed.FgColor())
	goassert.New(t, "Hello with Black").Equal(FgBlack.Sprintf("Hello with Black"))
	goassert.New(t, "\033[31mHello with Red\033[0m").Equal(FgRed.Sprintf("Hello with Red"))
	goassert.New(t, "\033[32mHello with Green\033[0m").Equal(FgGreen.Sprintf("Hello with Green"))
	goassert.New(t, "\033[33mHello with Yellow\033[0m").Equal(FgYellow.Sprintf("Hello with Yellow"))
	goassert.New(t, "\033[34mHello with Blue\033[0m").Equal(FgBlue.Sprintf("Hello with Blue"))
	goassert.New(t, "\033[35mHello with Magenta\033[0m").Equal(FgMagenta.Sprintf("Hello with Magenta"))
	goassert.New(t, "\033[36mHello with Cyan\033[0m").Equal(FgCyan.Sprintf("Hello with Cyan"))
	goassert.New(t, "\033[37mHello with Gray\033[0m").Equal(FgGray.Sprintf("Hello with Gray"))
}

func TestDecoration(t *testing.T) {
	goassert.New(t, false).Equal(Normal.IsBold())
	goassert.New(t, false).Equal(Normal.IsUnderline())
	goassert.New(t, "Hello with Normal Black").Equal(Normal.Sprintf("Hello with Normal Black"))
	goassert.New(t, true).Equal(Bold.IsBold())
	goassert.New(t, false).Equal(Bold.IsUnderline())
	goassert.New(t, "\033[1mHello with Bolded Black\033[0m").Equal(Bold.Sprintf("Hello with Bolded Black"))
	goassert.New(t, false).Equal(Underline.IsBold())
	goassert.New(t, true).Equal(Underline.IsUnderline())
	goassert.New(t, "\033[4mHello with Underlined Black\033[0m").Equal(Underline.Sprintf("Hello with Underlined Black"))
}

func TestStyle(t *testing.T) {
	goassert.New(t, "\033[1m\033[31mHello with Bolded Red\033[0m").Equal(FgRed.Bold(true).Sprintf("Hello with Bolded Red"))
	goassert.New(t, "\033[31mHello with Unbolded Red\033[0m").Equal(FgRed.Bold(false).Sprintf("Hello with Unbolded Red"))
	goassert.New(t, "\033[4m\033[31mHello with Underlined Red\033[0m").Equal(FgRed.Underline(true).Sprintf("Hello with Underlined Red"))
	goassert.New(t, "\033[31mHello with Ununderlined Red\033[0m").Equal(FgRed.Underline(false).Sprintf("Hello with Ununderlined Red"))
	goassert.New(t, "\033[4m\033[35mHello with Underlined Magenta\033[0m").Equal(FgRed.Underline(true).SetFgColor(FgMagenta).Sprintf("Hello with Underlined Magenta"))
	var buf bytes.Buffer
	FgRed.Fprintf(&buf, "Hello with Red")
	goassert.New(t, []byte("Hello with Red")).Equal(buf.Bytes())
}
