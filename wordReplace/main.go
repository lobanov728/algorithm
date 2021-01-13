package main

import (
	"math"
	"strings"
)

func main() {

}

func WordReplace(word1 string, word2 string) bool {
	var rune1 = []rune(word1)
	var rune2 = []rune(word2)
	if math.Abs(float64( len(rune1) - len(rune2))) > 1 {
		return false;
	}

	if len(rune1) == len(rune2) {
		var needChanges int
		for i, _ := range rune1 {
			if rune1[i] != rune2[i] {
				needChanges++
			}
			if needChanges > 1 {
				return false
			}
		}
		return true
	}

	if len(rune1) > len(rune2) {
		rune1, rune2 = rune2, rune1
	}
	var matches int
	for _, ch := range rune1  {
		if strings.Contains(string(rune2), string(ch)) {
			matches++
		}
	}
	return matches == len(rune1)
}