package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	for row := 0; row < n; row++ {
		for col := row + 1; col < n; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
	for row := 0; row < n; row++ {
		for col := 0; col < n/2; col++ {
			matrix[row][col], matrix[row][n-1-col] = matrix[row][n-1-col], matrix[row][col]
		}
	}

}
