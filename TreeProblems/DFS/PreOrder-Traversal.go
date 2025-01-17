// 144. Binary Tree Preorder Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    
    var dfs func(node *TreeNode)
    var anw []int

    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }

        anw = append(anw, node.Val)
        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root)
    return anw
}
