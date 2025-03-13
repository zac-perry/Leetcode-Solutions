// 80. Remove Duplicates from Sorted Array II
// Here, i use a weird two pointer approach
// Basically, as we loop, i check right and right-1
// If they are not equal, reset the total count (to track 2 max occurences), move the right value into the index where left is, then increment
// Otherwise, if they are equal, and count <= 1, then we can move and iterate as well

// This works since we are always checking the right (fast) pointer and what was before it, rather than in front
// By using left as a sort of "index marker" to know where to move the elemnt, and a count to track number of occurences (since this is sorted), it works.
/*
EXAMPLE: 
At i = 1: nums[1] == nums[0] → count = 1, nums[1] = 1, current = 2.
At i = 2: nums[2] == nums[1] → count = 2 → Skip.
At i = 3: nums[3] != nums[2] → count = 0, nums[2] = 2, current = 3.
At i = 4: nums[4] == nums[3] → count = 1, nums[3] = 2, current = 4.
At i = 5: nums[5] != nums[4] → count = 0, nums[4] = 3, current = 5

[1, 1, 1, 2, 2, 3]
 l. r
 // count++
 [1, 1, 1, 2, 2, 3]
     l. r -> count == 2 already, reset count = 0, move
 [1, 1, 1, 2, 2, 3]
        l. r -> replace, count + 1
[1, 1, 2, 2, 2, 3]
           l r -> same, count = 1, replace
[1, 1, 1, 2, 2, 3] 
             l. r -> diff, swap     
*/
func removeDuplicates(nums []int) int {
    if len(nums) == 1 {
        return 1
    } else if len(nums) == 0 {
        return 0
    }

    left, right := 1, 1
    count := 0

    for right < len(nums) {
        if nums[right] != nums[right-1] {
            count = 0
            nums[left] = nums[right]
            left++
        } else {
            count++
            if count <= 1 {
                nums[left] = nums[right]
                left++
            }
        }
        right++
    }
    return left
}
