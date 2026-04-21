package main

import (
	"fmt"
	"io"
	"os"

	"github.com/tibs/markdowncli/converter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: markdowncli <file.md>")
		fmt.Fprintln(os.Stderr, "       cat file.md | markdowncli -")
		os.Exit(1)
	}

	var input []byte
	var err error

	if os.Args[1] == "-" {
		input, err = io.ReadAll(os.Stdin)
	} else {
		input, err = os.ReadFile(os.Args[1])
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	result, err := converter.Convert(string(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting markdown: %v\n", err)
		os.Exit(1)
	}

	if result.Title != "" {
		fmt.Fprintf(os.Stderr, "Title: %s\n", result.Title)
	}
	fmt.Fprintf(os.Stderr, "Words: %d | Reading time: ~%d min\n", result.WordCount, result.ReadingTime)
	fmt.Print(result.HTML)
}
