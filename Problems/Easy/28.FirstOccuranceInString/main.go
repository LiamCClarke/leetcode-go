package main

import (
	"fmt"
)

func main() {
	//result := strStr("sadbutsad", "sad") // 0
	//result := strStr("hello", "ll") // 2
	result := strStr("mississippi", "issipi")
	fmt.Print(result)
}

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	for i, char := range haystack {
		if len(needle) > len(haystack)-i {
			return -1
		}
		if char == []rune(needle)[0] {
			starterIndex := i
			for j, needleChar := range needle {
				if needleChar != rune(haystack[i]) {
					break
				} else {
					if len(needle) == j+1 {
						return starterIndex
					} else {
						i++
					}
				}
			}
		}
	}
	return -1
}

// < -------- Solution 2 -------- >//
// Uses less memory
// `<=` means the for loop wont stretch out of bounds
func strStr2(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
