package main

import (
	"fmt"

	"github.com/daniilchuits/pre-trainee-practice/internal/wordcount"
)

func main() {
	s := "Go is a programming language. Go is expressive, concise, clean, and efficient.Concurrency is a core idea in Go programming. Programming in Go is fun."
	strs := wordcount.Tokens(s)

	for i, str := range strs {
		fmt.Printf("№%d: %s\n", i, str)
	}
}
