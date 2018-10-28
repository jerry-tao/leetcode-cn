func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var res,tmp *ListNode
    if l1==nil{
        return l2
    }
    if l2==nil{
        return l1
    }
    if l1.Val< l2.Val{
        res=l1
        l1 = l1.Next
    }else{
        res=l2
        l2 = l2.Next
    }
    tmp=res
    for l1!=nil && l2!=nil{
        if l1.Val < l2.Val{
            res.Next = l1
            l1 = l1.Next
        }else{
            res.Next = l2
            l2 = l2.Next
        }
        res = res.Next
    }

    for l1 != nil{
        res.Next = l1
        l1 = l1.Next
        res = res.Next
    }
        for l2 != nil{
        res.Next = l2
            l2=l2.Next
        res = res.Next
    }
    return tmp
}
