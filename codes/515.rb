# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Integer[]}
def largest_values(root)
    return [] if root.nil?
    result = []
    queue = []
    queue << root
    while(queue.size>0) do
      sum = nil
      temp = []
      while(queue.size>0) do
          node = queue.shift
          sum = node.val if  sum.nil? || node.val > sum
          temp << node.left if node.left
          temp << node.right if node.right
      end
      queue = temp
      result << sum
    end
    return result
end
