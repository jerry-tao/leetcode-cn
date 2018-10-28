# @param {Integer[]} nums
# @return {Integer}
def maximum_gap(nums)
    return 0 if nums.length < 2
    arrs = []
    nums = nums.sort
    nums.each_with_index do |i,item|
        arrs<< (i-nums[item+1]).abs unless nums[item+1]==nil
    end
    arrs.max
end
