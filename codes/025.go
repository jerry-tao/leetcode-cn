/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head==nil||k==1{return head}
    c:=&ListNode{Next: head}
    pre,cur,count:=c,head,0
    reverse:=func(p,n *ListNode)*ListNode{
        last := p.Next
        current := last.Next
        for current!=n{
            last.Next = current.Next
            current.Next = pre.Next
            pre.Next = current
            current = last.Next
        }
        return last
    }
    for cur!=nil{
        count++
        if count%k==0{
            pre = reverse(pre,cur.Next)
            cur = pre.Next
        }else{
            cur = cur.Next
        }
    }
    return c.Next
}
