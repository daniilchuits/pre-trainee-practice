package wordcount

var stopWords = map[string]struct{}{
	"a":   {},
	"an":  {},
	"the": {},
	"is":  {},
	"in":  {},
	"on":  {},
	"of":  {},
	"and": {},
	"or":  {},
	"to":  {},
	"for": {},
}

func isStopWord(word string, stopWords map[string]struct{}) bool {
	_, ok := stopWords[word]
	return ok
}
