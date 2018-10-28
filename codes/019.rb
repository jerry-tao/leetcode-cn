# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val)
#         @val = val
#         @next = nil
#     end
# end

# @param {ListNode} head
# @param {Integer} n
# @return {ListNode}
def remove_nth_from_end(head, n)
    arrs = []
    while head
      arrs << head.val
      head = head.next
    end
    arrs.delete_at(-n)
          l = nil
  v = nil
  arrs.each_with_index do |item, index|
    if index == 0
      l = ListNode.new(item)
      v = l.next = ListNode.new(0) unless index+1 == arrs.length
    else
      v.val=item
      v = v.next = ListNode.new(0) unless index+1 == arrs.length
    end
  end
  l


end
