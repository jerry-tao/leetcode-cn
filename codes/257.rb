# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {String[]}
def get_path(node, path, paths)
  paths  << "#{path}#{node.val}" unless (node.left || node.right)
  get_path(node.left, "#{path}#{node.val}->", paths) if node.left
  get_path(node.right, "#{path}#{node.val}->", paths) if node.right
end
def binary_tree_paths(root)
    paths = []
    get_path(root, '', paths) if root
    return paths
end
