# @param {Integer} n, a positive integer
# @return {Integer}
def hamming_weight(n)
    res=0
    while n!=0 do
        res+=1
        n &= (n-1)
    end
    return res
end
