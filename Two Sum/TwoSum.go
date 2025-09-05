package main

import "fmt"

func twoSum(nums []int, target int) [2]int {
	var sum int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]
			if sum == target {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}

func main() {
	var nums = [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	var target int = 5
	var index = [2]int{}
	index = twoSum(nums[:], target)
	fmt.Println(index)
}
