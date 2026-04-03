func productExceptSelf(nums []int) []int {
    numsLen := len(nums)
    
    result := make([]int, numsLen)
    left := 1
    for i := 0; i < numsLen; i++ {
        result[i] = left
        left *= nums[i] 
    }
    
    right := 1
    for i := numsLen-1; i >= 0; i-- {
        result[i] *= right
        right *= nums[i]
    }

    return result
}