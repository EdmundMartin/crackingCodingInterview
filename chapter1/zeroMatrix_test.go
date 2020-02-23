package chapter1

import "testing"

func TestZeroMatrix(t *testing.T) {
	testMatrix := [][]int{
		[]int{0, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 1, 1},
	}
	result := ZeroMatrix(testMatrix)
	if result[0][1] != 0 && result[0][2] != 0 {
		t.Errorf("did not propogate zeros")
	}
	if result[1][0] != 0 && result[2][0] != 0 {
		t.Errorf("did not propogate zeros")
	}
}