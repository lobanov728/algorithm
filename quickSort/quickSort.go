package algorithm

import (
	"fmt"
)

func getPivot(input []int) (int, int) {
	pos := len(input) - 1
	return input[pos], pos
}

func QuickSort(input []int) []int {
	fmt.Println("input", input)
	if len(input) < 2 {
		return input
	}
	pivot, pos := getPivot(input)
	fmt.Println("pivot", pivot)
	var greater, lesser []int
	for _, elem := range append(input[0:pos], input[(pos + 1):]...) {
		if elem >= pivot {
			greater = append(greater, elem)
		}
		if (elem < pivot) {
			lesser = append(lesser, elem)
		}
	}
	fmt.Println("greater", greater)
	fmt.Println("lesser", lesser)
	return append(append(QuickSort(lesser), pivot), QuickSort(greater)...)
}

