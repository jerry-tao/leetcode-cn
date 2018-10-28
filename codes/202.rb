# @param {Integer} n
# @return {Boolean}
def is_happy(n)
    return true if n ==1
    return false if n == 4
    begin
    # when use begin rescue it will be faster. WTF
    new_n = 0
    n.to_s.split('').each do |i|
      new_n += i.to_i*i.to_i
    end
    is_happy(new_n)
    rescue SystemStackError
    return false
    end
end
