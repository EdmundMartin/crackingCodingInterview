package chapter1

import "testing"

type indexPoint struct {
	row int
	col int
	value int
}

func TestRotateMatrix(t *testing.T) {
	matrix := [][]int{
		[]int{1, 0, 2},
		[]int{0, 0, 0},
		[]int{3, 0, 4},
	}
	result := RotateMatrix(matrix)
	points := []indexPoint{
		{row: 0, col: 0, value: 3},
		{row: 0, col: 2, value: 1},
		{row: 2, col: 0, value: 4},
		{row: 2, col: 2, value: 2},
	}
	for _, point := range points {
		if result[point.row][point.col] != point.value {
			t.Errorf("expected row %d, col %d to be equal to %d", point.row, point.col, point.value)
		}
	}
}