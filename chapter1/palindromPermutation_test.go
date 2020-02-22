package chapter1

import "testing"

func TestIsPermutationOfPalindrome(t *testing.T) {
	example := "Tact Cao"
	result := IsPermutationOfPalindrome(example)
	if !result {
		t.Errorf("provided example is an example of a palindrome")
	}
	example = "Tactoid Cao"
	result = IsPermutationOfPalindrome(example)
	if result {
		t.Errorf("provided input is not example of a palindrome")
	}
}