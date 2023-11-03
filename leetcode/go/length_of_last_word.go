package main

import "strings"

func lengthOfLastWord(s string) int {
	splittedString := strings.Split(strings.TrimSpace(s), " ")
	return len(splittedString[len(splittedString)-1])
}
