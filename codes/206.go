/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var pre, next *ListNode
    for head!=nil{
        if head.Next==nil{
            head.Next = pre
            break
        }
        next = head.Next
        head.Next = pre
        pre = head
        head = next
    }
    return head
}
