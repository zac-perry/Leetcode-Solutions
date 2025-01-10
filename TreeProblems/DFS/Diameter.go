// 543. Diameter of Binary Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(left int, right int) int {
    if left > right {
        return left
    }
    return right
}

// approach here was post-order DFS
// basically traverse to each node, determine it's "height" with respect to its parent
func diameterOfBinaryTree(root *TreeNode) int {
    var dfs func (node *TreeNode) int
    diameter := 0

    dfs = func (node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := dfs(node.Left)
        right := dfs(node.Right)
        diameter = max(diameter, left + right)

        return 1 + max(left, right)
    }

    dfs(root)

    return diameter
}
