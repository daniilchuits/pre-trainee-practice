package pipeline

func FilterMin(in []int, min int) []int {
	switch {
	case in == nil:
		return nil
	case len(in) == 0:
		return []int{}
	}

	res := []int{}
	for i := range in {
		if in[i] >= min {
			res = append(res, in[i])
		}
	}
	return res
}
