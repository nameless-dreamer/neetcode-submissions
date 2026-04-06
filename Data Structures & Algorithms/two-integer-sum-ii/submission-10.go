func twoSum(numbers []int, target int) []int {
	numLen := len(numbers)

	low := 0
	high := numLen - 1

	for low < high {
		sum := numbers[low] + numbers[high]

		if target == sum {
			return []int{low + 1, high + 1}
		} else if target > sum {
			low++
		} else {
			high--
		}
	}

	return []int{}
}