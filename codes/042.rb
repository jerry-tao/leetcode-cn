# @param {Integer[]} height
# @return {Integer}
def trap(height)
    arr=[]
    i=0;res=0
    while i<height.size do
        if arr.size==0 || height[i]<=height[arr[0]]
            arr.unshift i
            i+=1
        else
            t = arr.shift
            next if arr.size==0
            res+= ([height[i],height[arr[0]]].min - height[t]) * (i-arr[0]-1)
        end
    end
    return res
end
