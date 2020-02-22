package chapter1

import "testing"

func TestUniqueChars(t *testing.T) {
	unique := "abcdefgh"
	result := UniqueChars(unique)
	if !result {
		t.Errorf("%s was expected to be unique", unique)
	}
	dupes := "abcdeghaa"
	result = UniqueChars(dupes)
	if result {
		t.Errorf("%s was expected to contain duplicates", dupes)
	}
}
