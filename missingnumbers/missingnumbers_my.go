package missingnumbers

func Missing(slice []int) []int {
	m := make(map[int]int)
	for _, n := range slice {
		m[n] = 1
	}

	res := make([]int, 0, 2)
	var i int
	for i < len(slice)+2 {
		if m[i+1] != 1 {
			res = append(res, i+1)
		}
		i++
	}
	return res
}
