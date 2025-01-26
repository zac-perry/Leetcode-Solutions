// 19. Remove nth Node from End of List
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // Start here is to just move the 'current' node until it is right before the correct index n
 // then, you can start moving a second, 'prev' node behind it until current = nil 
 // As a result, the prev will then equal the node right before the one that needs to be removed. Can then just unlink it
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    current := head
    prev := dummy
    
    for current != nil {
        if n > 0 {
            current = current.Next
            n--
        } else {
            current = current.Next
            prev = prev.Next
        }
    }

    prev.Next = prev.Next.Next
    return dummy.Next
}
