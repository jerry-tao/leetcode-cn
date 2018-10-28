# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @param {Integer} v
# @param {Integer} d
# @return {TreeNode}
def add_one_row(root, v, d)
    if d == 1
        tmp = TreeNode.new v
        tmp.left = root
        return tmp
    end
    level = 2
    dfs(root, level, v, d)
    return root
end

def dfs(node, level, v, d)
    return if node.nil?
    if level==d
            tmp = node.left
            node.left = TreeNode.new v
            node.left.left = tmp
            tmp = node.right
            node.right = TreeNode.new v
            node.right.right = tmp
    else
        dfs(node.left, level+1, v, d)
        dfs(node.right, level+1, v, d)
    end
end
