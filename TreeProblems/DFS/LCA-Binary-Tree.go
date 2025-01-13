// 236. Lowest Common Ancestor of a Binary Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // Here, i perform a post-order dfs traversal to find the target nodes and determine the LCA
 // Esentially, by returning whenveer a node is p or q, this helps search and find the targets. 
 // Then, i can check and see if either leftTree or RightTree found a target. If both did, return the current node. Otherwise, return the non-nil one
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var dfs func (node *TreeNode) *TreeNode

    dfs = func(node *TreeNode) *TreeNode{

       if node == nil {
        return nil 
       } 

       // Found one of the targets
       if node == p || node == q {
        return node
       }

       leftTree := dfs(node.Left)
       rightTree := dfs(node.Right)

       // if a target was found in both subtrees, return the current node
       if leftTree != nil && rightTree != nil {
        return node
       } 

       // if a target was only found in one of the subtrees
       if leftTree != nil {
        return leftTree
       }

       return rightTree
    }

   return dfs(root)
}
