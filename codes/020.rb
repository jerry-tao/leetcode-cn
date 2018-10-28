# @param {String} s
# @return {Boolean}
def is_valid(s)
    stack = []
    s.each_char do |c|
        if [ '(',  '{',  '[' ].include?(c)
            stack << (c)
        end
        if [')', '}', ']'].include?(c)
            case c
            when ')'
              if stack[-1] != '('
                return false
            else
                stack.pop
            end
            when '}'
            if stack[-1] != '{'
                                return false
            else
                stack.pop
            end
            when ']'
            if stack[-1] != '['
                return false
            else
                stack.pop
            end
            end
        end
    end
    stack.empty? ? true : false
end
