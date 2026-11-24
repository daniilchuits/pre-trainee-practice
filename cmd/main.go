package main

import (
	"fmt"

	"github.com/daniilchuits/pre-trainee-practice/internal/wordcount"
)

func main() {
	s := "Go is a programming language. Go is expressive, concise, clean, and efficient.Concurrency is a core idea in Go programming. Programming in Go is fun."

	tokens := wordcount.Tokens(s)

	freq := wordcount.Count(tokens)
	pairs := wordcount.TopN(freq, 100)
	words := wordcount.DropStop(pairs)
	fmt.Print(words)
}
