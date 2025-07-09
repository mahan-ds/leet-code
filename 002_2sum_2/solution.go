package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(numbers, target)
	fmt.Println(res)

}

func twoSum(numbers []int, target int) []int {
	var i int = 0
	var j int = len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if target > sum {
			i += 1
		}
		if target < sum {
			j -= 1
		}
		if target == sum {
			return []int{i+1, j+1}
		}
	}
	return []int{}

}
