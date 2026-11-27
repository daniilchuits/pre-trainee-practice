package main

import (
	"fmt"

	"github.com/daniilchuits/pre-trainee-practice/internal/fileproc"
)

func main() {

	results := fileproc.Process([]string{
		"testdata/fileproc/ok.txt",
		"testdata/fileproc/missing.txt",
		"testdata/fileproc/panic.txt",
		"testdata/fileproc/empty.txt",
		"testdata/fileproc/panic_lower.txt",
		"testdata/fileproc/panic_trim.txt",
		"testdata/fileproc/directory",
	})

	for i := range results {
		fmt.Printf("№%d | %v\n", i+1, results[i])
	}

	results = fileproc.Process(nil)
	for i := range results {
		fmt.Printf("№%d | %v\n", i+1, results[i])
	}
}
