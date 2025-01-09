// 98. Validate Binary Search Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // Here, i do a pre-order DFS traversal
 // I use a min and max bounds
 // for each node, i ensure that min < node.Val < max
 // Doing it this way works in that each left node uses the previous nodes value as the max
 // Each right node uses the previous nodes value as the min 
 // This way, I can continually check if whatever node I am at is within the correct bounds.
func isValidBST(root *TreeNode) bool {
    var dfs func (node *TreeNode, min, max int)
    anw := true

    // pre order traversal
    dfs = func(node *TreeNode, min, max int) {
        
        if node == nil {
            return
        }
        
        // if the current node is not in bounds
        if node.Val <= min || node.Val >= max {
            anw = false
            return
        }

        dfs(node.Left, min, node.Val)
        dfs(node.Right, node.Val, max)
    }

    // min and max are just the min and max ints allowed here
    dfs(root, math.MinInt, math.MaxInt)

    return anw  
}

/********* EXAMPLE **********
      5
     / \
    1   4
       / \
      3   6

first -> -inf < 5 < inf -> good
    Left: (min = -inf, max = 5)
    Right: (min = 5, max = inf)

first -> Left -> -inf < 1 < 5 -> good

first -> Right -> 5 < 4 < inf -> not good, return false
****************************/
