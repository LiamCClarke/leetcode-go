package main

import (
	"strconv"
)

// ---------- SOLUTION 1 ---------- //
// Convert int to an array and work through each rune and compare to the other side
// This would break if x = 0
func isPalindrome_solution1(x int) bool {
	runes := []rune(strconv.Itoa(x))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// ---------- SOLUTION 2 ---------- //
func isPalindrome_solution2(x int) bool {
	// Check for negative numbers
	if x < 0 {
		return false
	}

	var rev int = 0
	var num int = x

	// reverse int
	for num > 0 {
		rev = rev*10 + num%10 // Create space for new number + final number of num
		num = num / 10        // remove end digit we just used
	}

	// Check reverse matches original
	return rev == x
}
