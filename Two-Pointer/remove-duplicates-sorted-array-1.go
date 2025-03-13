// 26. Remove Duplicates from a sorted array
// Two pointer approach
// Basically, left and right will loop over nums
// If we find an instance where the two values at these indicies do NOT match, we will replace elements (IN PLACE) in nums

// For example
// 1 1 2
// l = 0, r = 1, index = 1 -> skip since they match
// l = 1, r = 2, index = 1 -> 1 != 2: nums[1] = 2, index++
// now: nums = [1,2] (removed the duplicate)


// 1 1 2 2 3 4 5 6
// index = 1, left = 0, right = 1
// 1 1 skip
// 1 2, nums = 1,2,2,2,3,4,5,6, index = 2
// 2 2 ski
// 2 3, nums = 1,2,3,2,3,4,5,6, index = 3
// 3,4, nums = 1,2,3,4,3,4,5,6, index = 4
// 4,5, nums = 1,2,3,4,5,4,5,6, index = 5
// 5,6, nums = 1,2,3,4,5,6,5,6, index = 6
// 6 unique numbers, first 6 numbers are sorted and in order
func removeDuplicates(nums []int) int {
   index := 1
   left := 0
   right := left + 1

    for right < len(nums) {
        if nums[left] == nums[right] {
            left++
            right++
        } else {
            nums[index] = nums[right]
            index++
            left++
            right++
        }
    }
   return index
}
