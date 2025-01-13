// 125. Valid Palindrome
func isPalindrome(s string) bool {
    // This is gross and there are better ways to do it lol 
    // check isDigit / isLetter
    formatString := strings.ReplaceAll(s, " ", "")
    formatString = strings.Map(func(r rune) rune {
        if (int(r) >= 33 && int(r) <= 47) || 
        (int(r) >= 58 && int(r) <= 64) ||
        (int(r) >= 91 && int(r) <= 96) ||
        (int(r) >= 123 && int(r) <= 126) {
            return -1  // -1 removes the character
        }
        return r
    }, formatString)
    formatString = strings.ToLower(formatString)

    left := 0
    right := len(formatString) - 1

    // two pointer, one at the beginning (left), one at the end (right)
    for left < right {
        if formatString[left] != formatString[right] {
            return false
        }
        left++
        right--
    }

    return true
}
