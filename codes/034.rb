def search_range(nums, target)
    s,e = -1,-1
    nums.each_with_index do |i,index|
        if i == target && s==-1
            s = index
        end
        if s!=-1 && i==target
            e = index
        end
    end
    if s!=-1 && e==-1
        e=s
    end
    [s,e]
end
