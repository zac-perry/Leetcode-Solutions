// 1325. Delete Leaves With a Given Value
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // Post order DFS traversal 
 // Whenever the node.Val == target & it's a leaf node, return nil (remove it)
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    var dfs func (node *TreeNode) *TreeNode

    dfs = func(node *TreeNode) *TreeNode{
        if node == nil {
            return nil
        }

        node.Left = dfs(node.Left)
        node.Right = dfs(node.Right)

        if node.Val == target && node.Left == nil && node.Right == nil {
            return nil
        }
        return node
    }

    return dfs(root)
}
