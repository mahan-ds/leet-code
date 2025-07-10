package main

import (
	"fmt"
)

func main() {
	res := longestPalindrome("babbbbad")
	fmt.Println(res)

}

func longestPalindrome(s string) string {
	max_length := ""
	for i := 0; i < len(s); i++ {
		j := i
		k := i
		for j >= 0 && k < len(s) && s[j] == s[k] {
			word := s[j : k+1]
			k++
			j--
			if len(word) > len(max_length) {
				max_length = word
			}
		}

		j = i
		k = i + 1
		for j >= 0 && k < len(s) && s[j] == s[k] {
			word := s[j : k+1]
			k++
			j--
			if len(word) > len(max_length) {
				max_length = word
			}
		}

	}
	return max_length

}
