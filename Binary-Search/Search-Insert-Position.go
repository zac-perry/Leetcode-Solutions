// 35. Search Insert Position
// Given a sorted array and a target value, return the index where the target is found.
// If not found, return the index where it would be inserted while maintaining the sort order.
//
// Example: nums = [1,3,5,6], target = 5 -> returns 2
//          nums = [1,3,5,6], target = 2 -> returns 1
//
// Algorithm uses binary search:
// - Compare middle element with target
// - If middle < target, search right half (left = mid + 1)
// - If middle > target, search left half (right = mid - 1)
// - If equal, we found the index
//
// When target isn't found (left > right):
// - 'left' points to the first element greater than target
// - This is exactly where target should be inserted to maintain sorting
// - Works for all cases: beginning, middle, or end of array
func searchInsert(nums []int, target int) int {
    left := 0
    right := len(nums) - 1

    for left <= right {
        mid := left + (right - left) / 2

        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            return mid
        }
    }

    return left
}
