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
def find_bottom_left_value(root)
    $result = root.val
    $cur_level = -1
    level = 0
    dfs(root, level)
    return $result
end

def dfs(node, level)
    return if node.nil?
    $cur_level, $result = level, node.val if level>$cur_level
    dfs(node.left, level+1)
    dfs(node.right, level+1)
end
