package floyd

func Triangle(rows int) [][]int {
	current := 1
	res := make([][]int, rows)
	for i := 0; i < rows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j < i; j++ {
			res[i][j] = current
			current++
		}
	}
	return res
}
