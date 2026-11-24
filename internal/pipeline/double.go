package pipeline

func Double(in []int) []int {
	switch {
	case in == nil:
		return nil
	case len(in) == 0:
		return []int{}
	}

	res := make([]int, len(in))
	for i := range in {
		res[i] = (in[i]) * 2
	}
	return res
}
