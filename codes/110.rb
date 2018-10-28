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

def is_balanced(root)
    return true if root.nil?
  right = depth(root.right)
  left = depth(root.left)
  return false if left<0 || right <0 || (left-right).abs > 1
  return true
end

def depth(root)
    return 0 if root.nil?
    left = depth(root.left)
    right = depth(root.right)
    return -1 if left<0 || right <0 || (left-right).abs > 1
    return [left, right].max+1
end
