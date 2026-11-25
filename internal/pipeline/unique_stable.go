package pipeline

func UniqueStable(in []int) []int {
	switch {
	case in == nil:
		return nil
	case len(in) == 0:
		return []int{}
	}

	seen := make(map[int]struct{})
	res := []int{}
	for i := range in {
		if _, ok := seen[in[i]]; ok {
			continue
		}
		seen[in[i]] = struct{}{}
		res = append(res, in[i])
	}
	return res
}
