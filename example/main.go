package main

import (
	"fmt"
	"os"

	"github.com/hiro4bbh/go-log"
)

func main() {
	if golog.IsTerminal(os.Stdout) {
		fmt.Println("You are in terminal")

		fmt.Println(golog.FgBlack.Sprintf("Hello with Black"))
		fmt.Println(golog.FgRed.Sprintf("Hello with Red"))
		fmt.Println(golog.FgGreen.Sprintf("Hello with Green"))
		fmt.Println(golog.FgYellow.Sprintf("Hello with Yellow"))
		fmt.Println(golog.FgBlue.Sprintf("Hello with Blue"))
		fmt.Println(golog.FgMagenta.Sprintf("Hello with Magenta"))
		fmt.Println(golog.FgCyan.Sprintf("Hello with Cyan"))
		fmt.Println(golog.FgGray.Sprintf("Hello with Gray"))

		fmt.Println(golog.Normal.Sprintf("Hello with Normal Black"))
		fmt.Println(golog.Bold.Sprintf("Hello with Bolded Black"))
		fmt.Println(golog.Underline.Sprintf("Hello with Underlined Black"))

		fmt.Println(golog.FgRed.Bold(true).Sprintf("Hello with Bolded Red"))
		fmt.Println(golog.FgRed.Bold(false).Sprintf("Hello with Unbolded Red"))
		fmt.Println(golog.FgRed.Underline(true).Sprintf("Hello with Underlined Red"))
		fmt.Println(golog.FgRed.Underline(false).Sprintf("Hello with Ununderlined Red"))
		fmt.Println(golog.FgRed.Underline(true).SetFgColor(golog.FgMagenta).Sprintf("Hello with Underlined Magenta"))
	} else {
		fmt.Println("You are not in terminal")
	}

	logger := golog.New(os.Stdout, nil)
	logger.Debugf("Hello from DEBUG level\n")
	logger.Infof("Hello from INFO level")
	logger.Warnf("Hello from WARN level")
	logger.Errorf("Hello from ERROR level")
}
