def length_of_longest_substring(s)
  i=j=m=c=0
  exists = {}
  while j< s.length do
    if exists[s[j]]
        while s[i]!=s[j]
            exists[s[i]] = false
            i+=1
        end
        c = j-i
        i+=1
    else
        c += 1
        exists[s[j]] = true
    end
    j+=1
    m = m > c ? m : c
  end

  return m
end
