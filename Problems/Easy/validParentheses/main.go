package main

import (
	"fmt"
)

func main() {
	//result := isValid("()")
	result := isValid("([])")
	//result := isValid("()[]{}")
	//result := isValid("([)]")

	fmt.Print(result)
}

func isValid(s string) bool {
	stack := []rune{}

	if len(s) < 2 {
		return false
	}
	for _, bracket := range s {
		stack = append(stack, bracket)
		var peek rune = stack[len(stack)-1] // top of stack

		if isOpen(peek) {
			continue
		} else if isClosed(peek) {
			if len(stack) >= 2 {
				if isPair(stack[len(stack)-2], peek) {
					stack = stack[:len(stack)-1]
					stack = stack[:len(stack)-1]
					continue
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}

func isOpen(bracket rune) bool {
	openingBrackets := [3]rune{'(', '{', '['}

	for _, value := range openingBrackets {
		if bracket == value {
			return true
		}
	}
	return false
}

func isClosed(bracket rune) bool {
	closingBrackets := [3]rune{')', '}', ']'}

	for _, value := range closingBrackets {
		if bracket == value {
			return true
		}
	}
	return false
}

func isPair(openBracket rune, closeBracket rune) bool {
	bracketPairs := map[rune]rune{'{': '}', '[': ']', '(': ')'}

	if closeBracket == bracketPairs[openBracket] {
		return true
	} else {
		return false
	}
}
