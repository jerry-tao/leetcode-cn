/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    m:=make(map[int]int)
    return check(root, m,k)
}

func check(root *TreeNode,m map[int]int,k int)bool{
    if root==nil{return false}
    if _,ok:=m[root.Val];ok{return true}
    m[k-root.Val] = 1
    return check(root.Left,m,k) || check(root.Right,m,k)
}
