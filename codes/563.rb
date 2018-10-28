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
def find_tilt(root)
    $result = 0
    traverse(root)
    return $result
end
def traverse(root)
    return 0 if root.nil?
    left, right = traverse(root.left),traverse(root.right)
    $result += (left-right).abs
    return left+right+root.val
end
