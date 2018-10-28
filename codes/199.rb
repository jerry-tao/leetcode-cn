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
def right_side_view(root)
    return [] if root.nil?
    result = []
    level = 0
    bfs(root, level, result)
    return result
end

def bfs(root, level, result)
    return if root.nil?
    if result[level].nil?
        result[level] = root.val
    end
    bfs(root.right, level+1, result)
    bfs(root.left, level+1, result)
end
