/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    if root==nil{return root}
    sum:=0
    pot(root,sum)
    return root
}
func pot(root *TreeNode, sum int)int{
    if root.Right!=nil{sum=pot(root.Right,sum)}
    sum+=root.Val
    root.Val=sum
    if root.Left!=nil{sum=pot(root.Left,sum)}
    return sum
}
