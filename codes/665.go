func checkPossibility(nums []int) bool {
    cnt := 0
    n := len(nums)
    for i:= 1;i < n;i++ {
            if (nums[i] < nums[i - 1]) {
                if cnt == 1{ return false}
                if i>1 && nums[i] <= nums[i-2]{
                    nums[i] = nums[i-1]
                }
                cnt++
            }
        }
    return true
}
