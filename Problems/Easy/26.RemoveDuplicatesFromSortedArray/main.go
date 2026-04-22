package main

import (
	"fmt"
)

func main() {
	result := removeDuplicates([]int{1, 2})
	fmt.Print(result)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var counter = 1
	var indexToReplace int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if indexToReplace == 0 {
				indexToReplace = i
			}
			continue
		} else if nums[i] > nums[i-1] {
			if indexToReplace != 0 {
				nums[indexToReplace] = nums[i]
				indexToReplace += 1
			}
			counter++
			continue
		}
	}
	//fmt.Print(nums)
	return counter
}
