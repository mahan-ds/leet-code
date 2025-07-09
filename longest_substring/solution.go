package main




func main() {
	lengthOfLongestSubstring("ali")
}



func lengthOfLongestSubstring(s string) int {
	hash_map := map[rune]int{}
	count := 0
	for idx, letter := range s {
		val, ok := hash_map[letter]
		if !ok {
			hash_map[letter] = idx
			count++
		} else {
			return val
		}

		
	}
	return  6
    
}