package main

import "fmt"



func main() {
	g := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(g)
}

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	last:= -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target{
			right = mid - 1
		}else{
			if nums[mid] == target{
				last = mid
			}
			left = mid + 1

		}
	}


	left = 0
	right = len(nums) - 1
	first := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target{
			left = mid + 1
		}else{
			if nums[mid] == target{
				first = mid
			}
			right = mid - 1

		}
	}
	return []int{first, last}
    
}