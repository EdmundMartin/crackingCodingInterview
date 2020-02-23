package chapter2

import (
	"testing"
)

func sameIntSlice(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for idx, val := range s1 {
		if val == s2[idx] {
			return false
		}
	}
	return true
}

func TestDeleteMiddleNode(t *testing.T) {
	ll := SliceToLinkedList([]int{1, 2, 24, 2, 3, 4, 5, 15, 13, 15, 13, 7})
	nodeToDelete := KthToLast(ll, 6)
	DeleteMiddleNode(nodeToDelete)
	result := IterateValues(ll)
	if !sameIntSlice(result, []int{1, 2, 24, 2, 3, 5, 15, 13, 15, 13, 7}) {
		t.Errorf("expected node was not removed")
	}
}
