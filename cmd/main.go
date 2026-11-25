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

	// for _, pair := range pairs {
	// 	fmt.Printf("Word: %s Count: %d\n", pair.Word, pair.Count)
	// }

	s, err = accesslog.TestDataText()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	// sum := accesslog.Summarize(s, 3)
	_ = s

	// cases := []string{
	// 	"203.0.113.10 GET /x 99x 1",
	// 	"",
	// }
	// for _, tt := range cases {
	// 	fmt.Println(accesslog.Summarize(tt, 10))
	// 	fmt.Println()
	// }

}
