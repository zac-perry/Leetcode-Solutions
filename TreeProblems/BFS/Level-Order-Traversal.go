// 102. Binary Tree Level Order Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS Traversal
// Approach here involved performing a BFS traversal using a queue
func levelOrder(root *TreeNode) [][]int {
   if root == nil { return nil}

   queue := []*TreeNode{root} 
   vals := [][]int{}

   // Iterate over the queue
   for len(queue) > 0 {
        // level is based on the size of the queue.
        levelVals := []int{}
        levelSize := len(queue)

        // Loop again to get the size of the current level in the tree.
        // Will pop off the front of the queue, append the value, and add it's children at the next level if it has any
        for i := 0; i < levelSize; i++ {
            currentNode := queue[0]
            queue = queue[1:]
            levelVals = append(levelVals, currentNode.Val)
            if (currentNode.Left != nil) {
                queue = append(queue, currentNode.Left)
            }
            if (currentNode.Right != nil) {
                queue = append(queue, currentNode.Right)
            }
        }

        vals = append(vals, levelVals)
    }

    return vals
}
