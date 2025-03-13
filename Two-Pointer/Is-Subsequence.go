// 392. Is Subsequence 
// Two pointer approach, loop through both strings at the same time
// when you encounter a match, increment both (want subsequence to be in the same order in t)
// otherwise, increment through t
// Ever a case where we find the subsequence and j_s == len(s) can return true early rather than looping through the rest of the string
// If both iterators do NOT reach their length, then false. 
func isSubsequence(s string, t string) bool {
   if len(s) == 0 {
    return true
   }

    i_t := 0
    j_s := 0
    for i_t < len(t) {
        if  j_s == len(s) {
            return true
        }
        if s[j_s] == t[i_t] {
            j_s++
            i_t++
        } else {
            i_t++
        }
    }

    if i_t == len(t) && j_s == len(s) {
        return true
    }

    return false
}
