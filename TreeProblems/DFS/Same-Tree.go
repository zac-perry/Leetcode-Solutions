// 100. Same Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // approach here was a pre-order DFS traversal
 // Basically, when traversing each tree, ensure that the node values are always the same
 // If they are also both null, thats fine, return true. But, if there is ever a time only 1 is null, return false
 // Ensure that each side traverses and returns true
func isSameTree(p *TreeNode, q *TreeNode) bool {
    var dfs func(node1 *TreeNode, node2*TreeNode) bool

    dfs = func(node1 *TreeNode, node2*TreeNode) bool {
        // Again, if both are nil, no problem
        // If only one is nil, that's a problem (not equal, one tree is missing a node on that side)
        if node1 == nil && node2 == nil {
            return true
        } else if node1 == nil || node2 == nil {
            return false
        }

        if node1.Val != node2.Val {
            return false
        }

       l := dfs(node1.Left, node2.Left)
       r := dfs(node1.Right, node2.Right)

       return l && r
    }
    

    return dfs(p, q)
}
