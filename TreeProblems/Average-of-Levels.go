// 637. Average of Levels in Binary Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // BFS Traversal
 // Approach involves traversing the tree using BFS to visit each level
 // Then, at each level, I add the values of each node to a slice, which I can then use to average them together and store them
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil { return nil }

    queue := []*TreeNode{root}
    vals := []float64{}

    for len(queue) > 0 {
        levelSize := len(queue)
        currentLevel := []float64{}

        for i := 0; i < levelSize; i++ {
            currentNode := queue[0]
            queue = queue[1:]
        
            // determine other values to throw on there
            currentLevel = append(currentLevel, float64(currentNode.Val))
    
            if currentNode.Left != nil {
                queue = append(queue, currentNode.Left)
            }
            if currentNode.Right != nil {
                queue = append(queue, currentNode.Right)
            }
        }


        sum := 0.0
        for _, value := range(currentLevel) {
            sum += value
        }

        vals = append(vals, (sum/float64(len(currentLevel))))
    }

    return vals
}
