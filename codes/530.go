/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    res := math.MaxInt32
    var pre *TreeNode
    var iot func(*TreeNode)
    iot = func (node *TreeNode){
        if node.Left!=nil{iot(node.Left)}
        if pre!=nil && node.Val-pre.Val<res{
            res = node.Val-pre.Val
        }
        pre = node
        if node.Right!=nil{iot(node.Right)}
    }
    iot(root)
    return res
}
