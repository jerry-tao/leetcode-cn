# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

class BSTIterator
    # @param {TreeNode} root
    def initialize(root)
        @stack = []
        while root do
          @stack.unshift(root)
          root = root.left
        end
    end

    # @return {Boolean}
    def has_next
        @stack.size!=0
    end

    # @return {Integer}
    def next
        top=@stack.shift
        res = top.val
        if top.right
          top = top.right
          while top
             @stack.unshift(top)
             top = top.left
          end
        end
        return res
    end
end

# Your BSTIterator will be called like this:
# i, v = BSTIterator.new(root), []
# while i.has_next()
#    v << i.next
# end
