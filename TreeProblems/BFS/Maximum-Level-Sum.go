// 1161. Maximum Level Sum of a Binary Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    queue := []*TreeNode{root}
    level_track := 1
    final_level := 1
    level_total := root.Val
    
    // Normal BFS
    for len(queue) > 0 {
        curr_len := len(queue)
        total := 0
        
        // Traverse each level at once.
        // Keep track of the total at each level
        for i := 0; i < curr_len; i++ {
            //dequeue
            node := queue[0]
            queue = queue[1:]

            total += node.Val
            // enqueue
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        // Track the larget total found and save the level
        if total > level_total {
            level_total = total
            final_level = level_track
        }

        level_track += 1
    }

    return final_level
}
