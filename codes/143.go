/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head==nil || head.Next==nil{
        return
    }
    var fast, slow *ListNode
    fast=head
    slow=head
    for ;fast.Next != nil && fast.Next.Next !=nil;{
        fast = fast.Next.Next
        slow=slow.Next
    }
    last:=slow.Next
    slow.Next = nil
    last = reverse(last)
    var next,next2 *ListNode
    for tmp:=head;;{
        if tmp==nil || last==nil {
            break
        }
        next = tmp.Next
        next2 = last.Next
        tmp.Next = last
        tmp.Next.Next = next
        tmp = next
        last = next2
    }

}
func reverse(node *ListNode) *ListNode{
    var pre,next *ListNode
    for ;;{
        if node.Next==nil{
            node.Next = pre
            break
        }
        next = node.Next
        node.Next = pre
        pre = node
        node = next
    }
    return node
}

func plist(head *ListNode){
        for tmp:=head;tmp!=nil;{
            fmt.Println(tmp)
            tmp = tmp.Next
    }
}
