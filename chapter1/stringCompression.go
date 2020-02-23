package chapter1

import (
	"fmt"
	"strings"
)

// String Compression: Implement

func CompressString(str string) string {
	chars := strings.Split(str, "")
	prev := chars[0]
	count := 1
	result := ""
	for i := 1; i < len(chars); i++ {
		if prev == chars[i] {
			count++
		} else {
			result = result + fmt.Sprintf("%s%d", prev, count)
			count = 1
			prev = chars[i]
		}
	}
	result = result + fmt.Sprintf("%s%d", prev, count)
	return result
}
