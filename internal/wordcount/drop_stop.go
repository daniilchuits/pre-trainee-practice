package wordcount

func DropStop(pairs []Pair) []Pair {

	res := make([]Pair, 0, len(pairs))

	for i := range pairs {
		if ok := isStopWord(pairs[i].Word, stopWords); !ok {
			res = append(res, pairs[i])
		}
	}
	return res
}
