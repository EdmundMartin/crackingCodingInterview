package chapter2

import (
	"testing"
)

func TestKthToLast(t *testing.T) {
	array := []int{1, 2, 24, 2, 3, 4, 5, 15, 13, 15, 13, 7}
	ll := SliceToLinkedList(array)
	result := KthToLast(ll, 3)
	if result.Value != 13 {
		t.Errorf("expected value to be 13, got %d", result.Value)
	}
}
