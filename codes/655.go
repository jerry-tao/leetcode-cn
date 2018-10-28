/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
    if root==nil{return nil}
    h:=height(root)
    w:=int(math.Pow(2,float64(h)))-1
    res:=make([][]string,h)
    ptree(root,&res,0,w-1,0,w)
    return res
}

func ptree(node *TreeNode,res *[][]string,l,r,h,w int){
    if node==nil{return}
    if (*res)[h] == nil{ (*res)[h] = make([]string,w) }
    (*res)[h][(l+r)/2] = strconv.Itoa(node.Val)
    ptree(node.Left, res, l, (l+r)/2-1, h+1,w)
    ptree(node.Right, res, (l+r)/2+1,r , h+1,w)
}

func height(root *TreeNode)int{
    if root==nil{
        return 0
    }
    return max(height(root.Left),height(root.Right))+1
}
func max(a,b int)int{
    if a>b{return a}
    return b
}
