# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} t
# @return {String}
def tree2str(t)
  $str = ''
  traverse(t)
  return $str
end

def traverse(t)
  return if t.nil?
  $str << "#{t.val}"
  $str << "()" if t.left.nil? && t.right
  $str << "(" if t.left
  traverse(t.left)
  $str << ")" if t.left
  $str << "(" if t.right
  traverse(t.right)
  $str << ")" if t.right
end
