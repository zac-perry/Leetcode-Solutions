// 114. Flatten Binary Tree to Linked List
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // approach here is very strange. While the order should be a pre-order traversal, you can actually accomplish this
 // much easier with a Reverse Post Order Traversal

 // POST ORDER TRAVERSAL -> involves: 
 // 1. Traverse left
 // 2. Traverse right
 // 3. Process
 // But, to produce the same output as a pre-order, we reverse the left and right order 
 // 1. Traverse right
 // 2. Traverse left
 // 3. Process
func flatten(root *TreeNode)  {
    var prev *TreeNode
    var dfs func(node *TreeNode)
    
    // Reverse Post order DFS
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }

        dfs(node.Right)
        dfs(node.Left)

        // Set the right link to the "previous" node
        // Set left to nil
        node.Right = prev
        node.Left = nil

        // prev is now the current node
        prev = node
    }

    dfs(root)
}

/******** EXAMPLE *********
      1
     / \
    2   5
   / \   \
  3   4   6

** starting on the right side ** 
6 -> 
    6.right = prev (nil)
    6.left = nil 
    prev = 6

5 -> 
    5.right = prev (6)
    5.left = nil
    prev = 5

4 ->
    4.right = prev (5)
    4.left = nil
    prev = 4

3 ->
    3.right = prev (4)
    3.left = nil
    prev = 3

2 -> 
    2.right = prev (3)
    2.left = nil
    prev = 2

1 ->
    1.right = prev (2)
    1.left = nil
    prev = 1
    DONE

Doing it this way (basically backwards) will ensure that nothing gets overwritten when re-assigning the right nodes
Since you process them last and go from right to left, much easier and, again, you won't lose anything in the process
**************************/
