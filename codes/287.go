func findDuplicate(nums []int) int {
    slow:=0
    fast:=0
    for ;;{
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast{
            fmt.Println(slow)
            fmt.Println(fast)
            break
        }
    }
			fast = 0
			for ;;{
			fast = nums[fast]
			slow = nums[slow]
			if slow==fast{
			break
			}
			}
    return nums[slow]
}
