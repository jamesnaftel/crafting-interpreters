package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/scanner"
)

func main() {

	file := flag.String("file", "", "source file name to execute")
	prompt := flag.Bool("prompt", false, "golox prompt")

	flag.Parse()

	if *prompt {
		fmt.Println("golox interactive prompt")
		runPrompt()
		os.Exit(0)

	}

	if *file != "" {
		runFile(*file)
		os.Exit(0)
	}
}

func runFile(srcPath string) {
	src := readSrcFile(srcPath)
	run(src)
}

func readSrcFile(srcPath string) string {

	src, err := ioutil.ReadFile(srcPath)
	if err != nil {
		fmt.Printf("Error reading src file: %v", err)
	}

	return string(src)
}

func runPrompt() {

	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Golox > ")
		hasLine := scan.Scan()
		if hasLine == true {
			line := scan.Text()
			if line == "" {
				fmt.Println("Bye!")
				break
			}

			run(line)
		}
	}
}

func run(src string) {
	var s scanner.Scanner
	s.Init(strings.NewReader(src))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Println(s.TokenText())
	}

}
