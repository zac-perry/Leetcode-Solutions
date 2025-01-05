// 701. Insert into a Binary Search Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Here, we search the BST using DFS to determine which sides of the tree need to be traversed based on the val 
// This helps in finding the correct position faster
// Whenever the node is null, it will then insert the value
  // This works because at this point, it has found which side of the tree and which node it will for sure be a child of
// Then, returns the original 'node', or the root in this case to show the new tree
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    var dfs func(node *TreeNode) *TreeNode

    dfs = func(node *TreeNode) *TreeNode {
        if node == nil {
            return &TreeNode{Val: val}
        }

        if val < node.Val {
            node.Left = dfs(node.Left)
        } else {
            node.Right = dfs(node.Right)
        }
        
        return node
    }

    return dfs(root)
}
