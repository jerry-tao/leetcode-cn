# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {Integer[]} nums
# @return {TreeNode}
def construct_maximum_binary_tree(nums)
    root = TreeNode.new nums.max
    left = nums.slice_before(root.val).to_a
    right = nums.slice_after(root.val).to_a
    build_left(root, left.first) if left.size>1
    build_right(root, right.last) if right.size>1
    return root
end

def build_left(node, nums)
    node.left = TreeNode.new nums.max
    left = nums.slice_before(node.left.val).to_a
    right = nums.slice_after(node.left.val).to_a
    build_left(node.left, left.first) if left.size>1
    build_right(node.left, right.last) if right.size>1
end
def build_right(node, nums)
    node.right = TreeNode.new nums.max
    left = nums.slice_before(node.right.val).to_a
    right = nums.slice_after(node.right.val).to_a
    build_left(node.right, left.first) if left.size>1
    build_right(node.right, right.last) if right.size>1
end
