// 15. 3Sum
// two pointer approach using a 'first' pointer as well
// this 'first' pointer is fixed at the beginning, with left starting right after it
// Hardest part here is the duplicate checking
// When moving the 'first' pointer, need to ensure i am not looking at a duplicate value
// Same for the left pointer, to prevent getting the same answers more than once
func threeSum(nums []int) [][]int {
    anw := [][]int{}
    sort.Ints(nums)
    
    // Looping is kinda weird but want to make sure the 'first' pointer will never go past left
    for first := 0; first < len(nums)-2; first++ {
        // skip duplicates for the first 'fixed' pointer 
        if first > 0 && nums[first] == nums[first-1] {
           continue
        }
        
        left := first + 1
        right := len(nums) - 1

        // 2 pointer example:
        // [ 1, 2, 3, 4]
        //   ^  ^     ^
        for left < right {
            // calculate total
            // Make sure to increment left and look for duplicates, skipping them
            total := nums[first] + nums[left] + nums[right]
            if total == 0 {
                anw = append(anw, []int{nums[first], nums[left], nums[right]})
                left++
                for left < right && nums[left] == nums[left-1] {
                    left++
                }
            } else if total < 0 {
                left++
            } else if total > 0 {
                right--
            }
        } 
    }
    return anw
}
