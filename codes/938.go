/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
    num:=0
    if root==nil{return 0}
    if root.Val >=L && root.Val<=R{
        num+=root.Val
    }
    if root.Val>L { num+= rangeSumBST(root.Left,L,R)}
    if root.Val<R { num+= rangeSumBST(root.Right,L,R) }
    return num
}
