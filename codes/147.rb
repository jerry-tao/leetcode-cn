# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val)
#         @val = val
#         @next = nil
#     end
# end

# @param {ListNode} head
# @return {ListNode}
def insertion_sort_list(head)
    nums = []
    while head
      nums<< head.val
      head=head.next
    end
              l = nil
  v = nil
  nums.sort!.each_with_index do |item, index|
    if index == 0
      l = ListNode.new(item)
      v = l.next = ListNode.new(0) unless index+1 == nums.length
    else
      v.val=item
      v = v.next = ListNode.new(0) unless index+1 == nums.length
    end
  end
  l
end
