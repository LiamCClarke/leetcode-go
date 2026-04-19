package main

import (
	"fmt"
)

func main() {
	//result := romanToInt("XIV") //14
	result := romanToInt("III") //3
	//result := romanToInt("MCMXCIV") //1994
	fmt.Print(result)
}

func romanToInt(s string) int {
	numerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	var num int
	slice := []rune(s)

	for i := 0; i < len(slice); i++ {
		var lookAhead rune
		// on last iteration make lookAhead nil
		if i == len(slice)-1 {
			lookAhead = 0
		} else {
			lookAhead = slice[i+1]
		}
		var currentNumeralInt int = numerals[slice[i]]
		var lookAheadNumeralInt int = numerals[lookAhead]
		// If current number is greater or equal to next int, no subtraction needed
		if currentNumeralInt >= lookAheadNumeralInt {
			num += currentNumeralInt
			// else, number but be subtracted from next, then skip loop iteration for next
		} else {
			num += lookAheadNumeralInt - currentNumeralInt
			i++ // skip an iteration
		}
	}
	return num
}
