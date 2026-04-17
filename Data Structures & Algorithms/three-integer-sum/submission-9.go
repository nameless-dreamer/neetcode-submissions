func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sortedArr := qSort(nums)

	for i, v := range sortedArr {
		left := i + 1
		right := len(sortedArr) - 1

		if i > 0 && sortedArr[i] == sortedArr[i-1] {
			continue
		}

		for left < right {
			sum := sortedArr[left] + sortedArr[right] + v

			if sum == 0 {
				result = append(result, []int{
					sortedArr[left], v, sortedArr[right],
				})
				left++
				right--

				for left < right && sortedArr[left] == sortedArr[left-1] {
					left++
				}
				for left < right && sortedArr[right] == sortedArr[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

func qSort(nums []int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return nums
	}

	p := nums[numsLen/2]
	left := make([]int, 0)
	mid := make([]int, 0)
	right := make([]int, 0)

	for _, v := range nums {
		if v > p {
			right = append(right, v)
		} else if v < p {
			left = append(left, v)
		} else {
			mid = append(mid, v)
		}
	}

	temp := append(qSort(left), mid...)
	return append(temp, qSort(right)...)
}