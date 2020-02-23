package chapter1

// Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, it's entire row and columns
// are set to 0

func ZeroMatrix(matrix [][]int) [][]int {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == 0 {
				rows[r] = true
				cols[c] = true
			}
		}
	}

	for row, val := range rows {
		if val {
			for i := 0; i < len(matrix[0]); i++ {
				matrix[row][i] = 0
			}
		}
	}
	for col, val := range cols {
		if val {
			for i := 0; i < len(matrix); i++ {
				matrix[i][col] = 0
			}
		}
	}
	return matrix
}