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
def find_second_minimum_value(root)
    $min, $smin = root.val, root.val
    dfs(root)
    $min == $smin ? -1 : $smin
end

def dfs(node)
  return if node.nil?
  if node.val < $min
    $smin = $min
    $min = node.val
  end
  $smin = node.val if node.val > $min && node.val < $smin
  $smin = node.val if $min == $smin && node.val > $min
  dfs(node.left)
  dfs(node.right)
end
