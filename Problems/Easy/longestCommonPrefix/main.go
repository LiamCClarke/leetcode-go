package main

func main() {
	result := longestCommonPrefix([]string{"a"})
	print(result)
}

func longestCommonPrefix(strs []string) string {
	var commonPrefix string = ""

	// This loop will run for one of the strings length, if its the shortest it works and if its the longest, we will break in time
	// now var i is the respective rune
	for i := 0; i < len(strs[0]); i++ {
		currentChar := []rune(strs[0])[i] // First string index character
		for _, str := range strs {
			if len(str)-1 < i {
				return commonPrefix
			} else if []rune(str)[i] == currentChar {
				continue
			} else {
				return commonPrefix
			}
		}
		commonPrefix += string(currentChar)
	}
	return commonPrefix
}
