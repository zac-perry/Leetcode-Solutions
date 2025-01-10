// 226. Invert Binary Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // Used a pre order traversal
 // basically swap the left and right, continue parsing
 // Can be done with other tree traversals as well, this just made most sense to me
func invertTree(root *TreeNode) *TreeNode {
    var dfs func (node *TreeNode)

    dfs = func(node *TreeNode) {
        if node == nil {
            return
       } 

        temp := node.Left
        node.Left = node.Right
        node.Right = temp

        dfs(node.Left)
        dfs(node.Right)
    }

    dfs(root)
    return root
}
