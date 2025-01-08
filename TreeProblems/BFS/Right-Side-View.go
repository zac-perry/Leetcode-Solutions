// 199. Binary Tree Right Side View
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    // if the tree is empty, return nil
    if root == nil { return nil }

    queue := []*TreeNode{root}
    values := []int{}

    // Loop over the nodes on the queue
    for len(queue) > 0 {
        /*
            Idea here is to process the entire, current level (all nodes on the queue)
            This way, I can determine the right most node that is viewable (last node processed)
        */
        queue_len := len(queue)
        for i := 0; i < queue_len; i++ {

            // Usual BFS stuff
            node := queue[0]
            queue = queue[1:]

            // NOTE: If we are on the last node, then this is the 'right most' and can be added
            if i == queue_len - 1{
                values = append(values, node.Val)
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return values
}


/********** EXAMPLE: **********

        1
       / \
      2   3
       \    \
        5    4

Level 1: I add the root.
- When processing this level, this node happens to be the first and last, so it is viewable from the right and is added.

Level 2: [2,3]
- Here, on the inner loop, I go node by node, dequeue each one and enqueuing their children. I find that 3 is the right-most, so it is added.

Level 3: [5, 4]
- Here, again, on the inner loop, I go node by node. I find that 4 is the right-most, so it is added.

Final: [1,3,4] are all viewable 'from the right'
/*******************************/
