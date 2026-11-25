package accesslog

import (
	"sort"
	"strings"
)

func Summarize(text string, topN int) Summary {

	var (
		sum Summary
	)
	pathsFreq := make(map[string]int)
	seenIPs := make(map[string]struct{})

	statuses := make(map[int]int)

	lines := strings.Split(text, "\n") // splits by newline-separater, so each line == log-line

	for i := 0; i < len(lines); i++ {
		line, ok := ParseLine(lines[i])
		if !ok {
			sum.Skipped++
			continue
		}
		var emptyLine Line
		if line == emptyLine {
			continue
		}

		sum.Parsed++

		if _, ok := seenIPs[line.IP]; !ok {
			sum.UniqueIPs++
			seenIPs[line.IP] = struct{}{}
		}
		statuses[line.Status]++
		pathsFreq[line.Path]++
	}

	sum.ByStatus = statuses
	if topN <= 0 {
		return sum
	}

	var paths []PathCount
	for k, v := range pathsFreq {
		paths = append(paths, PathCount{Path: k, Count: v})
	}

	sort.Slice(paths, func(i, j int) bool {
		ok := paths[i].Count > paths[j].Count
		if ok {
			return true
		}
		if paths[i].Count == paths[j].Count {
			if paths[i].Path < paths[j].Path {
				return true
			}
		}
		return false
	})

	topN = min(len(paths), topN)
	sum.TopPaths = paths[:topN]
	return sum
}
