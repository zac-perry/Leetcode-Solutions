// 680. Valid Palindrome II
// Same as palindrome 1, except you can remove a char. In this case, whenever they are not equal
// try removing the left and see if it's still a palindrome. Same with the right. If not, return false
func checkPalindrome(s string, left, right int) bool {
    for left < right {
        if s[left] != s[right] {
            return false
        }

        left++
        right--
    }

    return true
}

func validPalindrome(s string) bool {
    left := 0
    right := len(s) - 1

    for left < right {
        if s[left] != s[right] {
            leftCheck := checkPalindrome(s, left + 1, right) 
            rightCheck := checkPalindrome(s, left, right - 1) 
            return leftCheck || rightCheck
        }
        left++
        right--
    }

    return true
}
