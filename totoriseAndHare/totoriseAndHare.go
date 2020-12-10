package algorithm

func FindDuplicate(nums []int) int {
	totorise := nums[0]
	hare := nums[0]

	for {
		totorise = nums[totorise]
		hare = nums[nums[hare]]
		if totorise == hare {
			break
		}
	}

	ptr1 := nums[0]
	ptr2 := totorise
	for {
		if (ptr1 == ptr2) {
			break
		}
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}

	return ptr1
}