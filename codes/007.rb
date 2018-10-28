def reverse(x)
    return 0 if x > (2**31-1)
    x = x.to_s.reverse
    if x.index('-')!=nil
        x.delete!('-')
        x.prepend('-')
    end
    x.to_i.abs > (2**31-1) ? 0 : x.to_i
end
