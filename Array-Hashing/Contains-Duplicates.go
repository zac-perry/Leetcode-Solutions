// 217. Contains Duplicates
// Just throw everything in a map as you go. If already there, return true
func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)

    for i := 0; i < len(nums); i++ {
        if m[nums[i]] {
            return true
        } else {
            m[nums[i]] = true
        }
    } 
    return false
}
