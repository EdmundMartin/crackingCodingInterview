package chapter2

import (
	"fmt"
	"testing"
)

func TestDeleteDupesBuffer(t *testing.T) {
	array := []int{1, 2, 24, 2, 3, 4, 5, 15, 13, 15, 13, 7}
	ll := SliceToLinkedList(array)
	DeleteDupesBuffer(ll)
	found := IterateValues(ll)
	if len(found) != NewSetFromSlice(found).Len() {
		fmt.Printf("expected unique values to be the same size as list values")
	}
}
