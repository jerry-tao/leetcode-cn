/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {

    res:=[]int{}
    if root==nil{
        return res
    }
    return pt(root,res)
}
func pt(node *TreeNode,res []int)[]int{
    if node.Left!=nil{res = pt(node.Left,res)}
    if node.Right!=nil{res = pt(node.Right,res)}
    res = append(res,node.Val)
    return res
}
