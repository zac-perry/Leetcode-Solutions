// 145. Binary Tree Postorder Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    
    var dfs func(node *TreeNode)
    var anw []int

    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }

        dfs(node.Left)
        dfs(node.Right)
        anw = append(anw, node.Val)
    }

    dfs(root)
    return anw
}
