/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy:=&ListNode{ Next: head }
    tmp := head
    last := dummy
    for tmp!=nil{
        n := tmp.Next
        if n==nil{break}
        last.Next = n
        tmp.Next = n.Next
        n.Next = tmp
        last = tmp
        tmp = tmp.Next
    }
    return dummy.Next
}
