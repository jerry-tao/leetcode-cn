# @param {Integer[]} digits
# @return {Integer[]}
def plus_one(digits)
    if digits[-1] != 9
        digits[-1]+=1
        return digits
    end

    stop = false

    result = digits.reverse!.map do |i|
        r = 0
        if stop
            r=i
        else
            if i!=9
                stop=true
                r=i+1
            end
        end
        r
    end

    result.reverse!

    result[0]==0 ? result.unshift(1) : result
end
