# @param {Integer} n, a positive integer
# @return {Integer}
def reverse_bits(n)
    sprintf('%032d',n.to_s(2)).reverse.to_i(2)
end
