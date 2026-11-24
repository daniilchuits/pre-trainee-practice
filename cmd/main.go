package main

import (
	"github.com/daniilchuits/pre-trainee-practice/internal/wordcount"
)

func main() {
	s := "Go is a programming language. Go is expressive, concise, clean, and efficient.Concurrency is a core idea in Go programming. Programming in Go is fun."

	tokens := wordcount.Tokens(s)
	tokensWithoutStop := wordcount.DropStop(tokens)
	freq := wordcount.Count(tokensWithoutStop)
	pairs := wordcount.TopN(freq, 3)

	_ = pairs
}
