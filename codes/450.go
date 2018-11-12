/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deleteMin(node *TreeNode) *TreeNode{
    if node.Left == nil{ return node.Right }
    node.Left = deleteMin(node.Left)
    return node
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root==nil{return nil}
    if key>root.Val{
        root.Right = deleteNode(root.Right, key)
    }else if key<root.Val{
        root.Left = deleteNode(root.Left, key)
    }else{
        tmp:=root
        if root.Right==nil{return root.Left}
        if root.Left == nil {return root.Right}
        min := root.Right
        for min!=nil && min.Left!=nil{
            min = min.Left
        }
        root = min
        root.Right = deleteMin(tmp.Right)
        root.Left = tmp.Left
    }
    return root
}
