package chapter1

import "strings"

// Check Permutations: Given two strings, write a method to decide if one is a permutation of the other

func quicksort(target []string) []string {
	if len(target) == 0 {
		return []string{}
	}
	if len(target) == 1 {
		return target
	}
	pivot := target[0]
	smaller := []string{}
	larger := []string{}
	for i := 1; i < len(target); i++ {
		if target[i] < pivot {
			smaller = append(smaller, target[i])
		} else {
			larger = append(larger, target[i])
		}
	}
	smaller = quicksort(smaller)
	larger = quicksort(larger)
	smaller = append(smaller, pivot)
	smaller = append(smaller, larger...)
	return smaller
}

func sortString(str string) string {
	sorted := quicksort(strings.Split(str, ""))
	joined := ""
	for _, val := range sorted {
		joined += val
	}
	return joined
}


func PermutationsUsingSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if sortString(s) != sortString(t) {
		return false
	}
	return true
}

func PermutationCharCounts(s string, t string) bool  {
	if len(s) != len(t) {
		return false
	}
	charCounts := [128]int{}
	for _, ch := range s {
		charCounts[ch]++
	}

	for _, ch := range t {
		charCounts[ch]--
		if charCounts[ch] < 0 {
			return false
		}
	}
	return true
}