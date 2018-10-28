/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    res:=[]int{}
    res=mt(root,res)
    return res[k-1]
}
func mt(node *TreeNode, res []int)[]int{
    if node.Left!=nil{res=mt(node.Left,res)}
    res = append(res, node.Val)
    if node.Right!=nil{res=mt(node.Right,res)}
    return res
}
