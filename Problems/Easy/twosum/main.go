package twosum

// ---------- SOLUTION 1 ---------- //
// O(n^2)
func twoSum_solution1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				break
			} else if (nums[i] + nums[j]) == target {
				result := []int{i, j}
				return result
			}
		}
	}
	result := []int{0, 0}
	return result
}

// ---------- SOLUTION 2 ---------- //
/*
* Still O(n^2)
* Because loops indexes cant be the same, the second loop can be (first index) + 1
* It wont fail as long as there is a guaranteed answer. Now you dont need the check for matching indexes
 */
func twoSum_solution2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// ---------- SOLUTION 3 ---------- //
/*
* Apparently better time complexity can be achieved using Hashmaps. Learning them now.
* O(n)
 */
func twoSum_solution3(nums []int, target int) []int {
	m := make(map[int]int)
	// initialise the map from the array so we can search using O(1)
	for index, value := range nums {
		m[value] = index
	}
	for index, value := range nums {
		// Find the number we need in the map to reach the target
		complement := target - value
		// Check for the number in the map, if found and not the same index then return both
		if j, ok := m[complement]; ok && j != index {
			return []int{index, j}
		}
	}
	// compiler candy
	return []int{}
}
