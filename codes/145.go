/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root==nil{return nil}
    stack:=[]*TreeNode{root}
    res:=[]int{}
    var tmp *TreeNode
    for len(stack)!=0{
        root = stack[len(stack)-1]
        if (root.Left == nil && root.Right==nil) || (tmp!=nil && (root.Left==tmp||root.Right==tmp)) {
            res = append(res, root.Val)
            stack = stack[:len(stack)-1]
            tmp = root
        }else{
            if root.Right!=nil{stack=append(stack,root.Right)}
            if root.Left!=nil{stack = append(stack,root.Left)}
        }
    }
    return res
}
