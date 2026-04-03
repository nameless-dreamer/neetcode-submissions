func hasDuplicate(nums []int) bool {
    table := map[int]bool{}
    flag := false

    for i:=0; i<len(nums); i++ {
        if table[nums[i]] {
            flag = true
        } else {
            table[nums[i]] = true
        }
    }

    return flag
}
