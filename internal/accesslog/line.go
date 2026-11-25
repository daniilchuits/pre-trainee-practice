package accesslog

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

type PathCount struct {
	Path  string
	Count int
}

// example
// 203.0.113.10 GET /api/books 200 142
