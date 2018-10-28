/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    res1:=[]int{}
    res2:=[]int{}
    res1=t(root1, res1)
    res2=t(root2, res2)
    if len(res1)!=len(res2){return false}
    for i,v:=range res1{
        if v!=res2[i]{return false}
    }
    return true
}
func t(node *TreeNode, res []int)[]int{
    if node.Left==nil && node.Right==nil{ return append(res,node.Val)}
    if node.Left!=nil{ res = t(node.Left,res) }
    if node.Right!=nil{ res = t(node.Right,res)}
    return res
}
