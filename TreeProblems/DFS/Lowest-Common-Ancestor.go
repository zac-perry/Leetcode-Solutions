// 235. Lowest Common Ancestor of a Binary Search Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
// aproach here: I used a pre order dfs traversal to check a few different cases
// 1. If the current node value is between p and q, then we have found the common ancestor and can return
// 2. If the current node equals either of these two, then we return. This handles the case were one of them is the common ancestor
// 3 & 4. Then, decide which way to traverse (if val < p and q, right. Otherwise, left)
// TLDR: Just try to search for p or q and, if ever in an optimal spot or one of them is found, return
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	
    var dfs func(node *TreeNode) *TreeNode

    dfs = func(node *TreeNode) *TreeNode {
        if node == nil {
            return nil
        }

        // if the node found is a value between the two p and q, return that node
        // else, if the node found is equal to either p or q, return (since ancestor could be one of these)
        // else if both p and q > current node, traverse right
        // otherwise, traverse left
        if (node.Val > p.Val && node.Val < q.Val) || (node.Val < p.Val && node.Val > q.Val) {
            return node
        } else if node == p || node == q {
            return node
        } else if p.Val > node.Val && q.Val > node.Val {
            return dfs(node.Right)
        } else if p.Val < node.Val && q.Val < node.Val {
            return dfs(node.Left)
        }

        return nil
    }

    return dfs(root)
}
