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
# @return {Boolean}
def get_path(node, path, sum)
  path.push node.val
  flag = false
  flag =  true if node.left==nil && node.right==nil && node.val == sum
  flag ||= get_path(node.left, path , sum-node.val) if node.left
  flag ||= get_path(node.right, path, sum-node.val) if node.right
  path.pop
  return flag
end
def has_path_sum(root, sum)
  return false unless root
    return get_path(root, [], sum)
end
