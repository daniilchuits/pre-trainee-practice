package pipeline

func Double(in []int) []int {
	switch {
	case in == nil:
		return nil
	case len(in) == 0:
		return []int{}
	}

	res := []int{}
	for i := range in {
		res = append(res, in[i]*2)
	}
	return res
}
