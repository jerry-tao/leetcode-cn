/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
    m:=make(map[int]int)
    find(root,m,0)
    res:=[]int{}
    max:=0
    for k,v:=range m{
        if v>max{
            max = v
            res = []int{k}
        }else if v==max{
            res = append(res,k)
        }
    }
    return res
}

func find(node *TreeNode, m map[int]int,sum int)int{
    if node==nil{return 0}
    sum=find(node.Left,m,sum)+find(node.Right,m,sum)+node.Val
    m[sum]+=1
    return sum
}
