// 42. Trapping Rain Water
// Two pointer approach (this one is great because it's O(n) and O(1) space)
// From both ends, determine the max height of the walls
// If the height on the right is greater, move left
    // Once done, find the amount of water you can add
        // Basically, the max left value is the absolute most water you can add since it is < right
            // limits the max water you can have
        // So, if maxLeft = 3 and height[left] == 3, cant add anything here (another wall)
        // But, if maxLeft = 3 and height[left] = 2, can add 1 to the water
    // similarly: 
        // If maxLeft > maxRight, then you are limited in height by the right side
        // Can just determine how much you can hold here (maxRight - height[right])
            // Again, same thing. if equal, then its another wall right in front and can't hold anything
            // If maxRight = 3 and height[right] = 1, then water += 2
            // move right--
    // Doing it this way will find all of the places you will have water based on the supposed max left and right heights found so far. 
    // NOTE: if maxLeft == maxRight, I just move the left value. 
func trap(height []int) int {
    left := 0
    right := len(height) - 1
    maxLeft := 0
    maxRight := 0
    water := 0

    for left < right {
        // get new maxes
        if height[left] > maxLeft {
            maxLeft = height[left]
        }
        if height[right] > maxRight {
            maxRight = height[right]
        }

        // move left++
        // Determine if you can have any water here based on the current value and maxLeft
        if maxLeft < maxRight {
            water += maxLeft - height[left]
            left++
        } else {
            // similar, move right -- 
            // determine if any water can be added based on the current right value and maxRight
            water += maxRight - height[right]
            right--
        }
    }

    return water
}
