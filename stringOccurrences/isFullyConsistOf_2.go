package stringOccurrences

import "fmt"

func isFullyConsistOf_2(word string, elems []string) bool {
	input := []rune(word)
	var currentFlag = true
	var nextFlag = false
	var calcNow bool
	var check1 bool
	var check2 bool
	for ptr := 0; ptr <= len(input); ptr++ {
		fmt.Println("ptr", ptr)
		calcNow = currentFlag
		if !currentFlag && !nextFlag {
			return false
		}
		check1 = false
		check2 = false
		for _, elem := range elems {
			if elem == string(input[ptr:ptr+1]) {
				check1 = true
			}
			if ptr + 1 <= len(input) && elem == string(input[ptr:ptr+2]) {
				check2 = true
			}
		}

		currentFlag = nextFlag || (calcNow && check1)
		nextFlag = calcNow && check2
	}
	return true
}