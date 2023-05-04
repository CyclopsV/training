package ArrayDiff

func ArrayDiff(a, b []int) []int {
	if len(a)==0 || len(b)==0 {
		return a
	}
	res := []int{}

	for _, va := range a {
		check := false
		for _, vb := range b {
			if va == vb {
				check = true
				break
			}
		}
		if ! check {
			res = append(res, va)
		}
	}

	return res
}
