// 383. Ransom note
func canConstruct(ransomNote string, magazine string) bool {
    // NOTE: I used a map here but could just use a slice instead 
    // In this case, make a slice of size 26 and add to each index based on the letter
    m := make(map[rune]int)

    for _, c := range(magazine) {
        m[c]++
    }    

    for _, c := range(ransomNote) {
        if m[c] == 0 {
            return false
        }
        m[c] -= 1
    }

    return true
}
