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
# @return {Integer}
def dfs(root,total, sum)
    count = 0
    count +=1 if root.val+total == sum
    count += dfs(root.left, total+root.val, sum) if root.left
    count += dfs(root.right, total+root.val, sum) if root.right
    return count
end
def path_sum(root, sum)
   return 0 if root.nil?
   dfs(root, 0, sum) + path_sum(root.left, sum) + path_sum(root.right, sum)
end
