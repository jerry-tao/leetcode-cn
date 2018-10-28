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
def diameter_of_binary_tree(root)
    $depth = 0
    depth(root)
    return $depth
end

def depth(root)
    return 0 if root.nil?
    right = depth(root.right)
    left = depth(root.left)
    $depth = [right+left, $depth].max
    return [left, right].max + 1
end
