// 1768. Merge Strings Alternately
// solution is a little overcomplicated but relies on checking 4 different cases
// basically, uses the index value (even or odd) to determine which char from which word to add
// Then, also checks cases in which one has had everything added and the other has not (add rest to the end) 
func mergeAlternately(word1 string, word2 string) string {
    endLength := len(word1) + len(word2)
    merged := make([]byte, endLength)

    index1 := 0
    index2 := 0

    for i := 0; i < endLength; i++ {
        // 1. If i is even and i haven't added all chars from word1.
        // 2. If i is odd and i haven't added all chars from word2.
        // 3. If i have added everything from word1, but not 2, add rest.
        // 4. If i have added everything from word2, but not 1, add rest.
        if i % 2 == 0 && index1 != len(word1){
            merged[i] = word1[index1]
            index1++
        } else if i % 2 != 0 && index2 != len(word2){
            merged[i] = word2[index2]
            index2++
        } else if index1 == len(word1) && index2 != len(word2){
            merged[i] = word2[index2]
            index2++
        } else {
            merged[i] = word1[index1]
            index1++
        }
    }

    return string(merged)
}

/* SIMPLER APPROACH */ 
/* Basically, loop and add chars in word1 and word2 simultaneously until one is done.
// Then, loop through them just to confirm you have everything else added at the end
i := 0
    j := 0
    merged := "" 

    for i < len(word1) && j < len(word2) {
        merged += string(word1[i]) + string(word2[j])
        i++
        j++
    }

    for i < len(word1) {
        merged += string(word1[i])
        i++
    }

    for j < len(word2) {
        merged += string(word2[j])
        j++
    }

    return merged
*
