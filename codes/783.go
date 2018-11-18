/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
    res:=math.MaxInt32
    var inorder func(*TreeNode)
    var pre *TreeNode
    inorder = func(node *TreeNode){
        if node.Left!=nil{ inorder(node.Left) }
        if pre==nil { pre=node }else{
            if node.Val-pre.Val<res { res = node.Val-pre.Val}
            pre = node
        }
        if node.Right!=nil{ inorder(node.Right) }
    }
    inorder(root)
    return res
}
