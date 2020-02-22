package chapter1
// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you can't use an
// additional data structure

import "strings"

func UniqueChars(str string) bool {
	if len(str) > 128 {
		return false
	}
	charSet := make(map[string]interface{})
	chars := strings.Split(str, "")
	for _, char := range chars {
		_, ok := charSet[char]
		if ok {
			return false
		}
		charSet[char] = nil
	}
	return true
}
