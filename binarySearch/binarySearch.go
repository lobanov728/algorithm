package main

import (
	"fmt"
)

func BinarySearch(needle float64, list []float64) float64 {
	low := 0
	high := len(list) - 1
	var mid int
	i := 0
	for low <= high {
		i++
		fmt.Println("step", i)
		mid = (low + high)/2
		fmt.Println(list[mid])
		if needle == list[mid] {
			fmt.Println("found", needle)
			return needle
		}
		if needle > list[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	panic("something wrong")
}

func main() {
	list := []float64{}

	for i := 1; i < 1000; i++ {
		list = append(list, float64(i))
	}
	BinarySearch(1, list)
}
