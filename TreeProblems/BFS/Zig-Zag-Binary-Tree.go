// 103. Binary Tree ZigZag Level Order Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil { return nil } 

    queue := []*TreeNode{root}
    vals := [][]int{}
    track := true

    // Looping over the queue
    for len(queue) > 0 {

        // Going to parse level by level
        curr_len := len(queue)
        curr_level := []int{}

        // normal BFS here, adding node values to curr_level
        for i := 0; i < curr_len; i++{
            node := queue[0]
            queue = queue[1:]

            curr_level = append(curr_level, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        // Using the track variable, I determine if the current level needs to be added in reverse or not
        // IF false, then reverse and add
        // Otherwise, just add
        if !track {
            reverse := []int{}
            for i := len(curr_level) - 1; i >= 0; i-- {
                reverse = append(reverse, curr_level[i])
            }
            vals = append(vals, reverse)
        } else {
            vals = append(vals, curr_level)
        }

        track = !track
    }

    return vals
}

/********** EXAMPLE *********

            3
          /   \
         9     20
              /  \
             15   7

output - [ [3], [20, 9], [15,7] ] 
****************************/
