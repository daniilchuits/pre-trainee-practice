package wordcount

func DropStop(words []string) []string {

	res := make([]string, 0, len(words))

	for i := range words {

		if _, ok := stopWords[words[i]]; !ok {
			res = append(res, words[i])
		}
	}
	return res
}
