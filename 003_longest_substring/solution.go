package main

import "fmt"

func main() {
	res := lengthOfLongestSubstring("alireab")
	fmt.Println(res)
}


func lengthOfLongestSubstring(s string) int {
	hash_map := map[rune]int{}
	pointer := 0
	max_length := 0
	for idx, letter := range s {
		value, ok := hash_map[letter]
		if ok && value >= pointer {
			pointer = value + 1
		}

		hash_map[letter] = idx
		length := idx - pointer + 1
		if length >= max_length {
			max_length = length
		}

	}
	return max_length

}