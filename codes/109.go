/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    tmp:=[]int{}
    for head!=nil{tmp = append(tmp, head.Val);head=head.Next}
    var bt func([]int) *TreeNode
    bt=func(vals []int) *TreeNode{
        if len(vals)==0{return nil}
        mid:=len(vals)/2
        node:=&TreeNode{Val:vals[mid]}
        node.Left = bt(vals[:mid])
        node.Right = bt(vals[mid+1:])
        return node
    }
    return bt(tmp)
}
