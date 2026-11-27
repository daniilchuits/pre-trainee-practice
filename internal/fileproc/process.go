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

	pathParts := strings.Split(path, "/")
	res.Path = pathParts[len(pathParts)-1]

	defer func() {
		if r := recover(); r != nil {
			res.Panicked = true
			res.Err = fmt.Errorf("Panic: %v", r)
		}
	}()

	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("%w: %s", NotExists, path)
		}
		res.Err = err
		return res
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		res.Err = err
		return res
	}

	lines := strings.Split(string(b), "\n")
	if len(lines) > 0 {
		firstLine := strings.TrimSpace(lines[0])
		if firstLine == invalidFirstLine { // exact math, not strings.HasPrefix
			panic("fileproc: boom")
		}
	}

	res.Bytes = len(b)

	res.Lines = bytes.Count(b, []byte{'\n'})

	if len(b) > 0 && b[len(b)-1] != '\n' {
		res.Lines++
	}

	return res
}
