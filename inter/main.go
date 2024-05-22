package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	chars := make(map[rune]bool)
	ans := make([]rune, 0, len(str1))

	for _, ch := range str1 {

		if _, oK := chars[ch]; !oK && containsRune(str2, ch) {
			chars[ch] = true
			ans = append(ans, ch)
		}
	}
	fmt.Println(string(ans))

}

func containsRune(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
