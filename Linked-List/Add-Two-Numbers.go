// 2. Add Two Number
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // Approach was to traverse both lists simultaneously, adding their values and calculating the carry
 // If there was a carry, it would be added to the next two values, since this is technically going in 'reverse', kinda like adding on paper
 // Additionally, i do everything in one loop, where if one list runs out of values, i continue traversing just the other to make sure i add everything
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    head := &ListNode{}
    current := head

    for l1 != nil || l2 != nil {
        //fmt.Println("ADDING: ", l1.Val, l2.Val)
        val := 0
        if l1 != nil && l2 != nil {
            val = l1.Val + l2.Val + carry
            l1 = l1.Next
            l2 = l2.Next
        } else if l1 != nil {
            val = l1.Val + carry
            l1 = l1.Next
        } else {
            val = l2.Val + carry
            l2 = l2.Next
        }
        carry = val / 10
        if val >= 10 {
            val = val % 10// problwem here
        }
        
        current.Val = val

        if l1 == nil && l2 == nil {
            break
        }
        current.Next = &ListNode{}
        current = current.Next
    }

    if carry != 0 {
        current.Next = &ListNode{}
        current = current.Next
        current.Val = carry
    }

    return head
}
