package main

import (
	"fmt"
	"os"

	"github.com/daniilchuits/pre-trainee-practice/internal/accesslog"
	"github.com/daniilchuits/pre-trainee-practice/internal/wordcount"
)

func main() {

	s, err := wordcount.TestDataText()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	tokens := wordcount.Tokens(s)
	tokensWithoutStop := wordcount.DropStop(tokens)
	freq := wordcount.Count(tokensWithoutStop)
	pairs := wordcount.TopN(freq, 3)

	_ = pairs

	line, ok := accesslog.ParseLine("203.0.113.10 GET /api/books 200 142")
	fmt.Print(line, ok)

}
