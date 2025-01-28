// 206. Reverse Linked List
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // Go through and just swap the links
 // make sure to keep a prev node
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    var node *ListNode = head

    for node != nil {
        temp := node.Next
        node.Next = prev
        prev = node
        node = temp
    }

    return prev
}
