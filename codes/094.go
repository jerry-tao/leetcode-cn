/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    stack := []*TreeNode{}
    res:=[]int{}

    for root!=nil || len(stack) != 0{
        if root.Left!=nil{
            stack = append(stack,root)
            root = root.Left
        }else{
            res = append(res,root.Val)
            root = root.Right
            for root==nil && len(stack) !=0{
                root = stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                res = append(res,root.Val)
                root = root.Right
            }
        }
    }
    return res
}
