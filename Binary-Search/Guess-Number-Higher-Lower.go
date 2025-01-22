/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

// 374. Guess Number Higher or Lower
// Here, I perform a binary search from the values 1 to n
// Based on the results from the guess() API function, i then move the left and right accordingly
func guessNumber(n int) int {
    left := 1
    right := n

    for left <= right {
        mid := left + (right - left) / 2
        guessValue := guess(mid)
        if guessValue == 0 {
            return mid
        } else if guessValue == -1 {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1
}
