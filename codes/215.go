func findKthLargest(nums []int, k int) int {
    sort.Ints(nums)
    l := len(nums)
    return nums[l-k]
}
