// 700. Search in a Binary Search Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Here, i use DFS to find the wanted node. In this case, when i find a node with a val == val, then i return it's subtree (just return the node)
func searchBST(root *TreeNode, val int) *TreeNode {
    var helperNode *TreeNode
    var dfs func(node *TreeNode)

    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }

        dfs(node.Left)
        if node.Val == val {
            helperNode = node
        }
        dfs(node.Right)

        return
    }

    dfs(root)

    return helperNode
}
