# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Boolean}
def check_symmetric(left, right)
  return true if left.nil? && right.nil?
  return false if left.nil? || right.nil?
  return false if left.val != right.val
  return check_symmetric(left.right, right.left) && check_symmetric(left.left, right.right)
end

def is_symmetric(root)
  return true if root.nil?
  return check_symmetric(root.left, root.right)
end
