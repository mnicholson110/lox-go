package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mnicholson110/lox-go/lox"
)

var hadError bool = false

func main() {
	switch len(os.Args[1:]) {
	case 0:
		runPrompt()
	case 1:
		runFile(os.Args[1])
	default:
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	}
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(74)
	}
	run(string(bytes))
	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {
	for {
		fmt.Println("Welcome to Lox! Press Ctrl+C to exit.")
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		run(line)
		hadError = false
	}
}

func run(source string) {
	scanner := lox.NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func LoxError(line int, message string) {
	lox.ErrorHandle(line, message)
	hadError = true
}