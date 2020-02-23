package chapter1

import "testing"

func TestOneEditAway(t *testing.T) {
	result := OneEditAway("google", "doogle")
	if !result {
		t.Error("expected edit distance of one")
	}
	result = OneEditAway("google", "doodle")
	if result {
		t.Error("edit distance of two")
	}
}