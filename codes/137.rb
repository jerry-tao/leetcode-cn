# @param {Integer[]} nums
# @return {Integer}
def single_number(nums)
    hash = Hash.new 0
    nums.each do |i|
        hash[i] += 1
    end
    return hash.select {|key, value| value==1}.keys.first
end
