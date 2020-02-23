package chapter1

import "testing"

func TestStringRotation(t *testing.T) {
	result := StringRotation("waterbottle", "bottlewater")
	if !result {
		t.Errorf("expected to be valid string rotation")
	}
	result = StringRotation("waterbottle", "dottlewater")
	if result {
		t.Errorf("expected to be invalid string rotations")
	}
}
