package main

import (
	"fmt"
	"os"

	"github.com/hiro4bbh/go-log"
)

func main() {
	promptStyle := golog.FgBlack
	if golog.IsTerminal(os.Stdout) {
		promptStyle = promptStyle.Bold(true)
	}
	term, err := golog.NewTerm(os.Stdin, promptStyle.Sprintf("> "), golog.TermConfig{
		History: true,
	})
	if err != nil {
		panic(err)
	}
	for {
		line, err := term.ReadLine()
		if err != nil {
			panic(err)
		} else if line == "" {
			break
		}
		fmt.Printf("you typed> %q\n", line)
	}
}
