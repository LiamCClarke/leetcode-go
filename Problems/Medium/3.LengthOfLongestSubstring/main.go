package main

import (
	"fmt"
	"strings"
)

func main() {
	//result := lengthOfLongestSubstring("dvdf")
	result := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(result)
}

// If I ever try this again, the more optimal solution is to store each char of s in a [byte]int array where int is the last known index of that char
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var longestSubstring string = s[0:1]
	var templongestSubstring string
	for _, char := range s {
		if strings.ContainsRune(templongestSubstring, char) {
			if len(templongestSubstring) > len(longestSubstring) {
				longestSubstring = templongestSubstring
			}
			templongestSubstring = templongestSubstring[strings.IndexRune(templongestSubstring, char)+1:] + string(char)
		} else {
			templongestSubstring += string(char)
			if len(templongestSubstring) > len(longestSubstring) {
				longestSubstring = templongestSubstring
			}
		}

	}
	return len(longestSubstring)
}
