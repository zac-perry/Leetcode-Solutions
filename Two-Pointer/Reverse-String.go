// 344. Reverse String
func reverseString(s []byte)  {
    left := 0
    right := len(s) - 1
    for left < right {
        temp := s[left]
        s[left] = s[right]
        s[right] = temp
        left++
        right--
    }
}
