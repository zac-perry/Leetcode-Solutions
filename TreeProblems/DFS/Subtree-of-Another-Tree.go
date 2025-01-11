// 572. Subtree of Another Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Approach here was very odd
// Esentially, i have two pre-order dfs traverals used for different reasons
// 1. I search until i find a node within tree1 that == a node in tree2
// 2. Once found, i then traverse the subtree to ensure that both are equal (l && r)
// 3. If found, i can return true
// 4. Otherwise, I keep searching for a potential starting point within tree 1. 
// 5. Then, i can return either l || r, since the subtree can be contained in either the left side or right side of the main tree
 func checkSubtree(node1 *TreeNode, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil {
        return true
    }

    if node1 == nil || node2 == nil {
        return false
    }
    
    l := false
    r := false

    if node1.Val == node2.Val {
        l = checkSubtree(node1.Left, node2.Left)
        r = checkSubtree(node1.Right, node2.Right)
    } else {
        return false
    }

    return l && r
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var searchDfs func(node1 *TreeNode, node2 *TreeNode) bool

    searchDfs = func(node1 *TreeNode, node2 *TreeNode) bool {
        if node1 == nil && node2 == nil {
            return true
        }

        if node1 == nil || node2 == nil {
            return false
        }

        l := false
        r := false
        val := false

        if node1.Val == node2.Val {
            val = checkSubtree(node1, node2)
        }

        if val {
            return true
        }

        l = searchDfs(node1.Left, node2)
        r = searchDfs(node1.Right, node2)

        return l || r
    }

    return searchDfs(root, subRoot)
}
