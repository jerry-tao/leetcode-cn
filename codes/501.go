/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    if root==nil{return nil}
    res:=[]int{}
    max:=0
    cur:=math.MinInt16
    cnt:=0
    return po(root,res,&max,&cur,&cnt)
}

func po(node *TreeNode, res []int,max,cur,cnt *int)[]int{
    if node.Left!=nil{res = po(node.Left,res,max,cur,cnt)}
    if *cur==node.Val{
        *cnt++
        if *cnt>*max{
            res = []int{*cur}
            *max = *cnt
        }else if *cnt==*max{
            res = append(res,*cur)
        }
    }else{
        *cur=node.Val
        *cnt=1
        if *cnt>*max{
            *max=*cnt
            res=[]int{*cur}
        }else if *cnt==*max{res=append(res,*cur)}
    }
    if node.Right!=nil{res = po(node.Right,res,max,cur,cnt)}
    return res
}
