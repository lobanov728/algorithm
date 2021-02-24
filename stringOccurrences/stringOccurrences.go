package stringOccurrences

var dict = []string{
	"a",
	"bc",
	"d",
	"ae",
	"abc",
	"de",
	"fg",
	"h",
	"ag",
	"ca",
	"clic",
	"cl",
	"c",
	"sc",
	"ar",
	"cd",
	"li",
	"fiance",
}

// find elements of dictionary occurrences in a words
func IsStringOccurrences(word string) bool {

	// map contains indexes of dictionary elements in word
	mapMatches := make(map[int][]int)
	for _, elem := range dict {
		indexes := getAllIndexes(word, elem)
		for _, index := range indexes{
			mapMatches[index] = append(mapMatches[index], len(elem))
		}
	}

	_, ok := mapMatches[0]
	if !ok {
		return false
	}

	pointer := 0
	totalLength := 0
	for i := 0; i < len(word); i++ {
		for _, count := range mapMatches[pointer] {
			if totalLength+count == len(word) {
				return true
			}
			_, ok := mapMatches[pointer+count]
			if ok {
				totalLength += count
				pointer += count
			}
		}
	}
	return false
}

func getAllIndexes(word string, substr string) []int {
	wordContent := []rune(word)
	var result []int
	for i := 0; i < len(wordContent); i++ {
		if string(wordContent[i:i+len(substr)]) == substr {
			result = append(result, i)
		}
	}

	return result
}
