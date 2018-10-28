# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Integer[][]}
def order_node(node, level, result)
    return if node.nil?
    result[level] = Array(result[level]) << node.val
    order_node(node.left, level+1, result)
    order_node(node.right, level+1, result)
end
def level_order(root)
    return [] if root.nil?
    level = 0
    result = []
    result[level] = Array(result[level]) << root.val
    order_node(root.left, level+1, result)
    order_node(root.right, level+1, result)
    return result
end
