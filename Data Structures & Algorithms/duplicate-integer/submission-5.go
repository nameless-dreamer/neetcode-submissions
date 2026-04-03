func hasDuplicate(nums []int) bool {
    numsLen := len(nums)
    table := make(map[int]bool, numsLen)
    flag := false

    for i := 0; i < len(nums); i++ {
        if table[nums[i]] {
            flag = true
        } else {
            table[nums[i]] = true
        }
    }

    return flag
}
