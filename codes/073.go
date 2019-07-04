func setZeroes(matrix [][]int) {
	is, js := make([]int, len(matrix)), make([]int, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				is[i] = 1
				js[j] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if is[i] == 1 || js[j] == 1 {
				matrix[i][j] = 0

			}
		}
	}

}
