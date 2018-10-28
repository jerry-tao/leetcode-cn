# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val)
#         @val = val
#         @next = nil
#     end
# end

# @param {ListNode} head
# @param {Integer} k
# @return {ListNode}
def reverse_k_group(head, k)
    return [] if head==nil
    arrs = []
    while head
      arrs << head.val
      head = head.next
    end


    arrs = arrs.each_slice(k).to_a
    if arrs[-1].length != k
        t = arrs.pop
    end
    arrs = arrs.map(&:reverse).flatten+Array(t)


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
