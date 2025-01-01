// 94. Binary Tree Inorder Traversal 
/**
* Definition for a binary tree node
* type TreeNode struct {
*   Val int
*   Left *TreeNode
*   Right *TreeNode
* }
**/

/******** Recursive Solution ********/ 
func helperFunction(node *TreeNode, vals *[]int) {
  // If the current node is not the root, then recurse the left, append, then right sides of the tree
  if node != nil {
    helperFunction(node.Left, vals)
    *vals = append(*vals, node.Val)
    helperFunction(node.Right, vals)
  }
}

func inorderTraversal(root *TreeNode) []int {
  vals := make([]int, 0)
  helperFunctin(root, &vals)
  return vals
}
