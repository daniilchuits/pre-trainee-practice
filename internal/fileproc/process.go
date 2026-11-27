package fileproc

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Process(paths []string) []Result {

	results := make([]Result, 0, len(paths))
	for i := range paths {
		results = append(results, processFile(paths[i]))
	}
	return results
}

func processFile(path string) Result {

	res := Result{}
	defer func() {
		if r := recover(); r != nil {
			res.Panicked = true
			res.Err = fmt.Errorf("Panic: %v", r)
		}
	}()

	f, err := os.Open(path)
	if err != nil {
		return Result{Err: err}
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return Result{Err: err}
	}

	res.Bytes = len(b)

	lines := strings.Split(string(b), "\n")

	res.Lines = len(lines)

	if len(lines) > 0 {
		firstLine := strings.TrimSpace(lines[0])
		if strings.HasPrefix(firstLine, "PANIC") {
			panic("fileproc: boom")
		}
	}
	return res
}
