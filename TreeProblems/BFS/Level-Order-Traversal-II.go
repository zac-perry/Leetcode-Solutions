// 107. Binary Tree Level Order Traversal II
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // Normal, level-order bfs traversal
 // Just making sure to prepend rather than append
func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }

    queue := []*TreeNode{root}
    vals := [][]int{}
   
    for len(queue) > 0 {
        currLen := len(queue)
        levelVals := []int{}
        for i := 0; i < currLen; i++ {
            node := queue[0]
            queue = queue[1:]

            levelVals = append(levelVals, node.Val)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }

            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        // prepend the slice to the front of vals
        vals = append([][]int{levelVals}, vals...)
    }

    return vals
}
