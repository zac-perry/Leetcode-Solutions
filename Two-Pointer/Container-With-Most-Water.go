// 11. Container with the Most Water
// Two pointer approach -> Start at either end
    // Go through and check the current heights (each is limited by the minimum value)
    // If one is less than the other, increment/decrement that side
    // Determine the area -> if new max, save
func maxArea(height []int) int {
    area := 0 
    left := 0
    right := len(height) - 1

    for left < right {
        currLen := right - left
        currHeight := 0
        if height[left] > height[right] {
            currHeight = height[right]
            right--
        } else if height[right] > height[left] {
            currHeight = height[left]
            left++
        } else if height[left] == height[right] {
            currHeight = height[left]
            left++
            right--
        }
        // get the area
        if currLen * currHeight > area {
            area = currLen * currHeight
        }
    }
    
    return area
}
