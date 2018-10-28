# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @param {Integer} sum
# @return {Integer[][]}
def get_path(node, path, paths, sum)
  path.push node.val
  paths  << path.clone if node.left==nil && node.right==nil && node.val == sum
  get_path(node.left, path , paths, sum-node.val) if node.left
  get_path(node.right, path, paths, sum-node.val) if node.right
  path.pop
end
def path_sum(root, sum)
    paths = []
    get_path(root, [], paths, sum) if root
    return paths
end
