package chapter1

import "strings"

func StringRotation(s1, s2 string) bool {
	length := len(s1)
	if length == len(s2) && length > 0 {
		joined := s1 + s1
		return strings.Contains(joined, s2)
	}
	return false
}