package main

import "fmt"




func main() {
	array := []int{1, 2, 3, 4}
	res := twoSum(array, 6)
	fmt.Println(res)
}


func twoSum(nums []int, target int) []int {
	hash_map := map[int]int{}
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if val, ok := hash_map[diff]; ok {
			return []int{val, i}
		}else {
			hash_map[nums[i]] = i
		}
	}
	return nil
    
}