# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Void} Do not return anything, modify root in-place instead.
def flatten(root)
    return if root.nil?
    left = root.left
    right = root.right
    root.left = nil
    flatten(left)
    flatten(right)
    root.right = left
    cur = root
    while(cur.right) do
        cur = cur.right
    end
    cur.right = right

end
