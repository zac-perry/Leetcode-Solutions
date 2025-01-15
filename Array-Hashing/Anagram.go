// 242. Valid Anagram
// here, i decided to use a singular hash map.
// Assuming both string were the same length, i could then traverse both at the same time
// Then, in the map, increment and decrement based on the strings found. Goal is to have a value of 0 for each key, showing that both contain the same chars
// So, when looking at the original string s, i increment, indicating that these are values found.
// Then, looking at the other string, i decrement indicating that the same value was also found

// NOTE: An easier, more intuitive approach, in my opinion, would be using 2 maps and checking if they are equal using DeepEqual, but that's just me.
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    m := make(map[rune]int)

    for i := 0; i < len(s); i++ {
        sRune := rune(s[i])
        tRune := rune(t[i])
        m[sRune]++
        m[tRune]--
    }

    for _, val := range m {
        if val != 0 {
            return false
        }
    }

    return true
}
