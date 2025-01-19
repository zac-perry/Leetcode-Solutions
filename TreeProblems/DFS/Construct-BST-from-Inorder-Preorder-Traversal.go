// 105. Construct binary tree from preorder and inorder traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    var dfs func([]int, []int, int) *TreeNode
    m := make(map[int]int)

    // store inorder values and their index for fast lookup
    for i, val := range(inorder) {
        m[val] = i
    }

    // weird sorta dfs preorder type thing goin on
    dfs = func(preorder []int, inorder []int, start int) *TreeNode {
        if len(preorder) == 0 || len(inorder) == 0 {
            return nil
        }

        root := &TreeNode{Val: preorder[0]}
        mid := m[root.Val] - start

        root.Left = dfs(preorder[1:mid+1], inorder[:mid], start)
        root.Right = dfs(preorder[mid+1:], inorder[mid+1:], start + mid + 1)

        return root
    }

    return dfs(preorder, inorder, 0)
}
// approach here is to use both the preorder and inorder information to create the tree.
// more specifically, using the preorder slice to get the root node values
// use the inorder slice to traverse the left + right subtrees correctly
// In dfs:
/*
    1. I get the current root (preorder[0], guarenteed each time)
    2. Then, i calculate the current mid -> this corresponds to the inorder traversal, as it seperates at the root value, giving you the left and right indicies. 
        - example (in order slice): [ 9, | 3, | 15, 20, 7]
                                    left, root,  right
                    
                    [15, 20, 7] -> in order slice at some point
                    [20, 15, 7] -> pre order slice
                    20 will be used as the root, allowing us to split the inorder slice to get the left and right sides (index is used as the middle)
    3. Then, I can traverse the subtrees 

    NOTE: I use a 'start' variable here to make sure I have the correct corresponding index in the slices for calculating the middle. This was needed so nothing went out of bounds
*/

/* EXAMPLE
ITERATION 1
pre - [3,9,20,15,7]
in -  [9,3,15,20,7]
root = 3
mid = indexInInOrderSlice - Start = 1 - 0 = 1 
Left = (pre[1:mid+1], in[:mid], 0)              - [9] [9] 0
Right = (pre[mid+1:], in[mid+1:], 1 + 0 + 1)    - [20, 15, 7] [15, 20, 7] 2

ITERATION 2 -> LEFT SIDE [9] [9] 0
pre - [9]
in - [9]
root = 9
mid = 0 - 0 = 0
Left = (pre[1:mid+1], in[:mid], 0)              - [] [] 0 -> RETURNS nil 
Right = (pre[mid+1:], in[mid+1:], 0 + 0 + 1)    - [] [] 1 -> RETURNS nil

ITERATION 3 -> RIGHT SIDE [20, 15, 7] [15, 20, 7] 2
pre - [20, 15, 7] 
in - [15, 20, 7]
root = 20 
mid = 3 - 2 = 1
Left = (pre[1:mid+1], in[:mid], 1)              - [15] [15] 1
Righ = (pre[mid+1:], in[mid+1:], 1 + 2 + 1)     - [7] [7] 4

ITERATION 4 -> LEFT SIDE [15] [15] 1
pre - [15]
in - [15]
root = 15
mid = 2 - 1 = 1
Left = (empty, empty, 1)     -> RETURNS nil
Right = (empty, empty, 3)    -> RETURNS nil 


ITERATION 5 -> RIGHT SIDE [7] [7] 4
pre - [7]
in - [7]
root = 7
mid = 0
Left = (empty, empty, 4)    -> RETURNS nil
Right = (empty, empty, 5)   -> RETURNS nil 

Final: all nodes and subtrees have been parsed.
*
