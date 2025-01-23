// 74. Search a 2D matrix
// Modified binary search
// here, i do a pretty standard binary search
// except, since this is a 2d matrix, you need to view it as a 'flat' list of elements
// to do this, you need to calculate the row and column within the 2d matrix based on the calculated middle element
    // row = i / n
    // col = i % n
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)        // rows
    n := len(matrix[0])     // columns
    left := 0
    right := (m * n) - 1    // 2d right calculation (can also do some weird len[len] type stuff)  

    for left <= right {
        mid := left + (right - left) / 2
        row := mid / n
        column := mid % n

        if matrix[row][column] == target {
            return true
        } else if matrix[row][column] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    
    return false
}
