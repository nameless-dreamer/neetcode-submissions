func maxArea(heights []int) int {
    leftBar := 0
    rightBar := len(heights)-1
    maxAmWater := 0
    
    for leftBar < rightBar {
        currentAmWater := 0
        currentGap := rightBar - leftBar

        if heights[leftBar] > heights[rightBar] {
            currentAmWater = heights[rightBar] * currentGap
            rightBar--
        } else {
            currentAmWater = heights[leftBar] * currentGap
            leftBar++
        }

        if currentAmWater > maxAmWater {
            maxAmWater = currentAmWater
        }
    }

    return maxAmWater
}
