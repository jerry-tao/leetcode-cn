/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    if root==nil{return nil}
    res :=[]int{}
    res = iot(root, res)
    head:=&TreeNode{Val: res[0]}
    tmp := head
    for i:=1;i<len(res);i++{
        tmp.Right = &TreeNode{Val: res[i]}
        tmp = tmp.Right
    }
    return head
}

func iot(node *TreeNode,res []int) []int{
    if node.Left!=nil{res=iot(node.Left,res)}
    res = append(res,node.Val)
    if node.Right!=nil{res=iot(node.Right,res)}
    return res
}
