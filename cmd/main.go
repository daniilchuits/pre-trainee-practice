package main

import (
	"fmt"
	"os"

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

	for _, pair := range pairs {
		fmt.Println(pair)
	}

}
