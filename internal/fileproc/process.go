package fileproc

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

const invalidFirstLine = "PANIC"

func Process(paths []string) []Result {

	results := make([]Result, 0, len(paths))
	for i := range paths {
		results = append(results, processFile(paths[i]))
	}
	return results
}

func processFile(path string) (res Result) {

	res.Path = path

	defer func() {
		if r := recover(); r != nil {
			res.Panicked = true
			res.Err = fmt.Errorf("Panic: %v", r)
		}
	}()

	f, err := os.Open(path)
	if err != nil {
		res.Err = err
		return res
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		res.Err = err
		return res
	}

	res.Bytes = len(b)

	res.Lines = bytes.Count(b, []byte{'\n'})

	if len(b) > 0 && b[len(b)-1] != '\n' {
		res.Lines++
	}

	lines := strings.Split(string(b), "\n")
	if len(lines) > 0 {
		firstLine := strings.TrimSpace(lines[0])
		if firstLine == invalidFirstLine { // exact math, not strings.HasPrefix
			panic("fileproc: boom")
		}
	}
	return res
}
