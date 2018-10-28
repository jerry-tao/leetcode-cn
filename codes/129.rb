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
def sum_numbers(root)
    return sum(root, 0)
end

def sum(node, sum)
    return 0 if node.nil?
    return sum*10+node.val if(node.right.nil? && node.left.nil?)
    return sum(node.left, sum*10+node.val) + sum(node.right, sum*10+node.val)
end
