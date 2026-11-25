package accesslog

import (
	"strconv"
	"strings"
)

func ParseLine(s string) (Line, bool) {

	s = strings.TrimSpace(s)
	switch {
	case s == "":
		return Line{}, true
	case s[0] == '#':
		return Line{}, false
	}

	parts := strings.Fields(s)
	if len(parts) != 5 {
		return Line{}, false
	}

	status, err := strconv.Atoi(strings.TrimSpace(parts[3]))
	if err != nil {
		return Line{}, false
	}

	bytes, err := strconv.Atoi(strings.TrimSpace(parts[4]))
	if err != nil {
		return Line{}, false
	}

	path := strings.TrimSpace(parts[2])
	if path[0] != '/' {
		return Line{}, false
	}

	return Line{
		IP:     strings.TrimSpace(parts[0]),
		Method: strings.TrimSpace(parts[1]),
		Path:   path,
		Status: status,
		Bytes:  bytes,
	}, true
}
