func deleteDuplicates(head *ListNode) *ListNode {
    tmp := head
    for tmp!=nil && tmp.Next!=nil{
        if tmp.Val==tmp.Next.Val{
            tmp.Next = tmp.Next.Next
        }else{
          tmp = tmp.Next
        }

    }
    return head
}
