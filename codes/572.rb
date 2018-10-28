# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} s
# @param {TreeNode} t
# @return {Boolean}
def is_same_tree(p, q)
    return true if q.nil? && p.nil?
    return false if q.nil? || p.nil?
    return false if p.val != q.val
    return is_same_tree(p.left, q.left) && is_same_tree(p.right, q.right)
end
def is_subtree(s, t)
  return true if is_same_tree(s, t)
  return false if s.nil? && !t.nil?
  return is_subtree(s.left, t) || is_subtree(s.right,t)
end
