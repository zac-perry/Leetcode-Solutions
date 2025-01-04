// 450. Delete Node in BST
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    When deleting a node, there are three cases we want to consider.
    1. When the target node has no children, can just return null
    2. When the target node has one child, you can just return the child
    3. When the target node has two children...
        - Get it's right successor
        - Find the minimum successor of this successor (Left)
        - If found, then replace the current (target) node value with the successor value
        - Remove the original successor from the subtree (Right)
    
    Example: key = 3
              5
             / \
            3   6
            / \   \
            2   4   7
                / \
            3.5 4.5
    
    - Here, it will find 3
    - Since 3 has two children, it will then go right (4) 
    - Then it will find the minimum child of 4 (Left, 3.5)
    - Once found, it will replace the original node (3) with the new child value (3.5) 
    - Then, it will delete 3.5 from it's original sub tree (calling dfs on Right)

    Resulting Tree
              5
             / \
            3.5 6
            /  \  \
            2   4  7
                 \
                  4.5
*/

func deleteNode(root *TreeNode, key int) *TreeNode {
    var dfs func(node *TreeNode, key int) *TreeNode

    dfs = func(node *TreeNode, key int) *TreeNode {
        if node == nil {
            return nil
        }

        // find which side to search
        if key < node.Val {
            node.Left = dfs(node.Left, key)
            return node
        } 
        if key > node.Val {
            node.Right = dfs(node.Right, key)
            return node
        } 
       
        // Found, deleting
        // Case 1.
        if node.Left == nil && node.Right == nil {
            return nil
        }

        // Case 2.
        if node.Left == nil {
            return node.Right
        } 
        if node.Right == nil {
            return node.Left
        }

        // Case 3. 
        successor := node.Right
        for successor.Left != nil {
            successor = successor.Left
        }

        // Reassign, then delete original child
        node.Val = successor.Val
        node.Right= dfs(node.Right, successor.Val)
    
        return node
    }

    return dfs(root, key)
}
