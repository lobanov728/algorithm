package stringOccurrences

import "fmt"

func IsFullyConsistOf(word string, elems []string, maxElemSize int) bool {
	input := []rune(word)
	
	elemSizes := make([]int, maxElemSize)
	for i := 0; i < maxElemSize; i++ {
		elemSizes[i] = i + 1;
	}
	fmt.Println("elem sizes", elemSizes)
	flags := make([]bool, maxElemSize)
	flags[0] = true // для того чтобы выполнить первую итерацию
	

	ptr := 0
	for ; ptr <= len(input); {
		fmt.Println("ptr", ptr)
		fmt.Println("word", string(input[ptr:]))
		fmt.Println("flags1", flags)
		contFlag := flags[0]; // если хотя одно совпадение есть
		returnFlag := contFlag; // если хотя бы есть один true в flags не выходим из цикла
		for i := 1; i < len(flags); i++ {
			flags[i - 1] = flags[i];
			returnFlag = returnFlag || flags[i];
		}
		flags[len(flags)-1] = false
		
		fmt.Println("flags2", flags)
		fmt.Println("returnFlag", returnFlag)
		fmt.Println("contFlag", contFlag)
		if !returnFlag {
			return false
		}
		if !contFlag {
			fmt.Println("cont")
			ptr++
			continue
		}

		for _, elemSize := range elemSizes {
			if ptr + elemSize <= len(input) {
				for _, elem := range elems {
					if elem == string(input[ptr:ptr+elemSize]) {
						flags[elemSize-1] = true
						break
					}
				}
			}
		}
		ptr++
	}
	return true
}
