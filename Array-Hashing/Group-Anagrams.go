// 49. Group Anagrams
// two approaches
// first uses an int slice of 26 as the key, representing all letters in the alphabet
// when looping through the strs, I can find which letters this string has and count them, using this as the key. Intuition here is that anagrams will have the same counts, so the key is easy to find
// Then, i append the string to the string slice val in the map

// Second approach is very similar. Except I reach each string as a []rune, then sort the []rune slice to get an alphabetical order of runes as my key
// Then, its basically the same thing. I just look to see if this key exists. If so, I append to the []string slice val. Otherwise, I insert. 
func groupAnagrams(strs []string) [][]string {
    // m := make(map[[26]int][]string)

    // for _, val := range strs {
    //     var count [26]int

    //     for _, char := range val {
    //         count[char - 'a']++
    //     }

    //     m[count] = append(m[count], val)
    // }

    // var anw [][]string
    // for _, val := range m {
    //     anw = append(anw, val)
    // }

    // return anw

    m := make(map[string][]string)
    for _, val := range strs {
        runes := []rune(val)
        sort.Slice(runes, func(i, j int) bool{
            return runes[i] < runes[j]
        })

        m[string(runes)] = append(m[string(runes)], val)
    }

    strings := [][]string{}
    for _, vals := range m {
        strings = append(strings, vals)
    }
    
    return strings
}
