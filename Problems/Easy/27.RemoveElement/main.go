package main

import (
	"fmt"
)

func main() {
	//result := removeElement([]int{3, 2, 2, 3}, 3)
	result := removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Print(result)
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var counter int
	var replaceIndex int = -1
	for i, num := range nums {
		if num == val {
			if replaceIndex == -1 {
				replaceIndex = i
			}
		} else {
			counter++
			if replaceIndex != -1 {
				nums[replaceIndex] = num
				nums[i] = val                    // replace current with val so it will be replaced if needed
				if nums[replaceIndex+1] == val { // If there are duplicate adjacent vals
					replaceIndex += 1
				} else {
					replaceIndex = -1 // reset index
				}
			}
		}
	}
	return counter
}
