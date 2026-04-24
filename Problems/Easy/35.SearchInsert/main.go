package main

func main() {

}

func searchInsert(nums []int, target int) int {
	for i, num := range nums {
		if target < nums[0] {
			return 0
		}
		if target == num {
			return i
		}
		if i == len(nums)-1 { // End of loop
			return i + 1
		}
		if (target > num) && (target <= nums[i+1]) {
			return i + 1
		}
	}
	return 0
}
