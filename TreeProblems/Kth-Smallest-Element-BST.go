// 230. Kth Smallest Element in a BST

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// helper function performs a DFS traversal on the tree
// We keep count of the nodes we visit. In this case, whenever the count == k, we return that value
// This works since it's looking for the kth-smallest value. So, whenever we are at the kth value, we just return
func helper(root *TreeNode, k int, counter *int, answer *int) {
    if root == nil {
        return
    }
    
    helper(root.Left, k, counter, answer)

    // Counting here. Once at the kth element, store and return
    *counter++
    if *counter == k {
        *answer = root.Val
        return
    }
    helper(root.Right, k, counter, answer)

    return
}

func kthSmallest(root *TreeNode, k int) int {
    counter := 0
    answer := 0

    helper(root, k, &counter, &answer)
   
    return answer
}
