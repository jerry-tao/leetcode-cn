def add_two_numbers(l1, l2)
   result = l1
  while l1 || l2
    l1.val = l1.val+l2.val

    if l1.val>9
      l1.val = l1.val%10
      l1.next ? l1.next.val+=1 : l1.next = ListNode.new(1)
    end

    l2.next = ListNode.new(0)  if l1.next && !l2.next
    l1.next = ListNode.new(0)  if l2.next && !l1.next

    l1 = l1.next
    l2 = l2.next
  end
  return result
end
