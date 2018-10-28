# @param {Integer[]} nums
# @return {String[]}
def summary_ranges(nums)
    result = []
    s,e = nil,nil
    nums.each_with_index do |i,index|
      s = i if s==nil
      if i+1 == nums[index+1]
          next
      else
          e = i
          result << [s,e]
          s = nil
      end
   end
   puts result.length
   result.map do |item|
       item[0] == item[1] ? item[0].to_s : item.join('->')
   end
end
