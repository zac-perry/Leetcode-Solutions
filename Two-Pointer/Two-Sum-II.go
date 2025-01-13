// 167. Two Sum II - Input Array Is Sorted
// did a generic two pointer approach -> binary search is probably better here though
func twoSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers) - 1

    for left < right {
        if numbers[left] + numbers[right] > target {
            right--
        } else if numbers[left] + numbers[right] < target {
            left++
        } else {
            return []int{left+1, right+1}
        }
    }

    return nil
}
