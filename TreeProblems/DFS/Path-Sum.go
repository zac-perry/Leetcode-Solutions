// 112. Path Sum
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // DFS pre-order traversal
 // Traverse and find the target sum assuming it is found on a leaf node
func hasPathSum(root *TreeNode, targetSum int) bool {
    
    var dfs func(node *TreeNode, currTotal int)
    found := false

    dfs = func(node *TreeNode, currTotal int) {
        if node == nil {
            return 
        }

        // Here, determine if the total is found and if it was also found at a leaf node
        currTotal = currTotal + node.Val
        if currTotal == targetSum && node.Left == nil && node.Right == nil {
            found = true
            return
        }

        dfs(node.Left, currTotal)
        dfs(node.Right, currTotal)
    }

    dfs(root, 0)
    return found
}
