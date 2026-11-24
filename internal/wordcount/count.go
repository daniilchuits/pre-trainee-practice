package wordcount

func Count(tokens []string) map[string]int {

	countMap := make(map[string]int)

	for i := range tokens {
		countMap[tokens[i]]++
	}
	return countMap
}
