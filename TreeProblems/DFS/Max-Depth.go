// 104. Max Depth of Binary Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // approach here: post order dfs traversal 
 // Get the height of the left and right sides, return the greater height + 1
func maxDepth(root *TreeNode) int {
    var dfs func(node *TreeNode) int

    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        leftH := dfs(node.Left)
        rightH := dfs(node.Right)

        if leftH > rightH {
            return leftH+1
        }
        return rightH + 1
    }

    return dfs(root)
}
