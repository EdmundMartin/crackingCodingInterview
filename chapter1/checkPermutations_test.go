package chapter1

import "testing"

func TestPermutationCharCounts(t *testing.T) {
	original := "cheburashka"
	permu := "burashkache"
	if !PermutationCharCounts(original, permu) {
		t.Errorf("expected to be a correct permutation")
	}

	nonPermu := "cheburuhkia"
	if PermutationCharCounts(original, nonPermu) {
		t.Errorf("expected not to be a permutation")
	}
}


func TestPermutationsUsingSort(t *testing.T) {
	original := "cheburashka"
	permu := "burashkache"
	if !PermutationsUsingSort(original, permu) {
		t.Errorf("expected to be a correct permutation")
	}

	nonPermu := "cheburuhkia"
	if PermutationsUsingSort(original, nonPermu) {
		t.Errorf("expected not to be a permutation")
	}
}