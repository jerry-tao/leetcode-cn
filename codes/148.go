/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head==nil || head.Next==nil{
        return head
    }
    p,f,s:=head,head,head
    for f!=nil && f.Next!=nil{
        p = s
        s=s.Next
        f=f.Next.Next
    }
    p.Next = nil
    return merge(sortList(head),sortList(s))
}
func merge(a,b *ListNode) *ListNode{
    if a==nil{return b}
    if b==nil{return a}
    if a.Val<b.Val{
        a.Next = merge(a.Next,b)
        return a
    }else{
        b.Next = merge(b.Next,a)
        return b
    }
}
