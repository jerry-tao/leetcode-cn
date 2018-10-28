# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Integer[]}
def inorder_traversal(root)
    result = []
    inorder(root, result)
    return result
end

def inorder(node, result)
    return if node.nil?
    inorder(node.left, result) if node.left
    result << node.val
    inorder(node.right, result) if node.right
end
