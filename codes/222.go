/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root==nil{return 0}
    lh,rh:=0,0
    l,r:=root,root
    for l!=nil{
        l=l.Left
        lh++
    }
    for r!=nil{
        r=r.Right
        rh++
    }
    if lh==rh{ return  int(math.Pow(float64(2), float64(lh))) - 1 }
    return countNodes(root.Left)+countNodes(root.Right)+1
}
