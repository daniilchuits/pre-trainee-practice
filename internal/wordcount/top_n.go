package wordcount

import (
	"sort"
)

type Pair struct {
	Word  string
	Count int
}

func TopN(freq map[string]int, n int) []Pair {

	if n <= 0 {
		return []Pair{}
	}

	res := make([]Pair, 0, len(freq))

	for k, v := range freq {
		res = append(res, Pair{
			Word:  k,
			Count: v,
		})
	}
	sort.Slice(res, func(i, j int) bool {
		if ok := res[i].Count > res[j].Count; ok {
			return true
		}

		if res[i].Count == res[j].Count {
			if res[i].Word < res[j].Word {
				return true
			}
			return false
		}
		return false
	})

	n = min(n, len(res))

	return res[:n]
}
