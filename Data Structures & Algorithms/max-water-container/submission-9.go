func maxArea(heights []int) int {
    currentGap := len(heights)-1

    leftBar := 0
    rightBar := len(heights)-1
    maxAmWater := 0
    
    for leftBar < rightBar {
        currentAmWater := 0

        if heights[leftBar] > heights[rightBar] {
            currentAmWater = heights[rightBar] * currentGap
            rightBar--
        } else if heights[leftBar] < heights[rightBar] {
            currentAmWater = heights[leftBar] * currentGap
            leftBar++
        } else {
            currentAmWater = heights[leftBar] * currentGap

			if heights[leftBar+1] < heights[rightBar-1] {
				rightBar--
			} else {
				leftBar++
			}
        }

        if currentAmWater > maxAmWater {
            maxAmWater = currentAmWater
        }
        currentGap--
    }

    return maxAmWater
}
