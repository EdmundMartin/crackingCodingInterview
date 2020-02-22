package chapter1

import "strings"

// Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or
// phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome
// does not need to be limited to just dictionary words. You can ignore casing and non-letter characters.
// EXAMPLE: Tact Coa OUTPUT: True (taco cat or acto cta etc.)

func getCharNum(r rune) int32 {
	if r < 97 || r > 122 {
		return -1
	}
	return r - 'a'
}

func characterFrequencyTable(phrase string) [26]int {
	charFreq := ['z'+1-'a']int{}
	for _, ch := range phrase {
		charNum := getCharNum(ch)
		if charNum != -1 {
			charFreq[charNum]++
		}
	}
	return charFreq
}

func checkMaxOneOdd(table [26]int) bool {
	foundOdd := false
	 for _, val := range table {
		if val % 2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	 return true
}


func IsPermutationOfPalindrome(phrase string) bool {
	phrase = strings.ToLower(phrase)
	table := characterFrequencyTable(phrase)
	return checkMaxOneOdd(table)
}