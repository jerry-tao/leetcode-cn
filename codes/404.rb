# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Integer}
def sum_of_left_leaves(root)
    result = 0
    return 0 if root.nil?
    result += root.left.val if root.left && root.left.left.nil? && root.left.right.nil?
    result += sum_of_left_leaves(root.left)
    result += sum_of_left_leaves(root.right)
    result
end
