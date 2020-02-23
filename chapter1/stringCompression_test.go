package chapter1

import "testing"

func TestCompressString(t *testing.T) {
	test := "aabcccccaaa"
	result := "a2b1c5a3"
	actualResult := CompressString(test)
	if result != actualResult {
		t.Errorf("expected %s got %s", result, actualResult)
	}
}
