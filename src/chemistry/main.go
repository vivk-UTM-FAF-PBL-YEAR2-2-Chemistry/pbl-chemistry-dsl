package main

import (
	"chemistry/config"
	"chemistry/interpreter"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		io.WriteString(os.Stderr, "The file was not found.\n")
		os.Exit(config.FILE_ERR)
	}

	file, err := ioutil.ReadFile(args[0])
	if err != nil {
		io.WriteString(os.Stderr, "The file was not found.\n")
		os.Exit(config.FILE_ERR)
	}

	content := string(file)
	lines := strings.Split(content, "\n")

	interpreter.Start(lines)
}
