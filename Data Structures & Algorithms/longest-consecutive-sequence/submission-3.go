func longestConsecutive(nums []int) int {
	numsTable := make(map[int]bool)
	for _, num := range nums {
		numsTable[num] = true
	}

	longest := 0
	for num := range numsTable {
		if !numsTable[num-1] {
            current := num
            length := 1

            for numsTable[current+1] {
                current++
                length++
            }


			if length > longest {
                longest = length
            }
		}
	}

	return longest	
}
