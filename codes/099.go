/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    res:=[]int{}
    nodes:=[]*TreeNode{}
    res,nodes=inorder(root,res,nodes)
    sort.Ints(res)
    for i,v:=range res{
        nodes[i].Val = v
    }
}
func inorder(node *TreeNode,res []int,nodes []*TreeNode) ([]int,[]*TreeNode){
    if node==nil{return res,nodes}
    res,nodes = inorder(node.Left,res,nodes)
    res = append(res, node.Val)
    nodes = append(nodes,node)
    res,nodes = inorder(node.Right,res,nodes)
    return res,nodes
}
