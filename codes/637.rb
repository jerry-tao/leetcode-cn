# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Float[]}
def average_of_levels(root)
    result = []
    queue = []
    queue << root
    while(queue.size>0) do
      sum = 0
        count = 0
      temp = []
      while(queue.size>0) do
          node = queue.shift
          sum += node.val
          count += 1
          temp << node.left if node.left
          temp << node.right if node.right
      end
      queue = temp
      result << sum*1.0/count
    end
    return result
end
