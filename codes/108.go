/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bst(nums []int,left,right int)*TreeNode{
    if (left > right) {return nil}
    mid:=(left+right)/2
    root:= &TreeNode{Val:nums[mid]}
    root.Left = bst(nums, left,mid-1)
    root.Right = bst(nums,mid+1,right)
    return root
}
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums)==0{return nil}
    return bst(nums, 0 , len(nums)-1);

}
