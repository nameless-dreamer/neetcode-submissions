func twoSum(nums []int, target int) []int {
    numsLen := len(nums)
    data := make(map[int]int, numsLen)
    
    for i := 0; i < numsLen; i++ {
        if value, k := data[target - nums[i]]; k {
            return []int{value, i}
        }
        data[nums[i]] = i
    }   
    return []int{}
}
