package accesslog

import (
	"fmt"
	"strings"
)

type Line struct {
	IP     string
	Method string
	Path   string
	Status int
	Bytes  int
}

type Summary struct {
	Parsed    int
	Skipped   int
	UniqueIPs int
	ByStatus  map[int]int
	TopPaths  []PathCount // up to N, count desc, path asc on ties
}

func (s Summary) String() string {

	var sb strings.Builder

	fmt.Fprintf(&sb, "Parsed=%d Skipped=%d UniqueIPs=%d\n",
		s.Parsed, s.Skipped, s.UniqueIPs,
	)

	fmt.Fprintln(&sb, "Statuses:")
	for status, count := range s.ByStatus {
		fmt.Fprintf(&sb, "Status:%d=%d\n", status, count)
	}
	fmt.Fprintf(&sb, "Top paths: %v with len=%d", s.TopPaths, len(s.TopPaths))
	return sb.String()
}

type PathCount struct {
	Path  string
	Count int
}

// example
// 203.0.113.10 GET /api/books 200 142
