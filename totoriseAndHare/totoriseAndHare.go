package main

import "fmt"

func main()  {
	// find duplicate numbers
	// given an array nums containing n + 1 integers where each integer is between 1 and n inclusive
	// prof that at least one duplicate numbers exists

	array := []int{
		10, 1, 3, 4, 5, 2, 3, 9, 7, 8, 6}
	fmt.Println(FindDuplicate(array))
}

func FindDuplicate(nums []int) int {
	totorise := nums[0]
	hare := nums[0]

	for {
		fmt.Println("Totorise", totorise)
		fmt.Println("Hare", hare)
		totorise = nums[totorise]
		hare = nums[nums[hare]]

		if totorise == hare {
			break
		}
	}

	ptr1 := nums[0]
	ptr2 := totorise
	for {
		fmt.Println("ptr1", ptr1)
		fmt.Println("ptr2", ptr2)
		if ptr1 == ptr2 {
			break
		}
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}

	return ptr1
}