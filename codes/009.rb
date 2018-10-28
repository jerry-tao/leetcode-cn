# @param {Integer} x
# @return {Boolean}
def is_palindrome(x)
  return false if x<0
  y = 0;  z = x
  while(x!=0) do y = y*10 + x%10;x = x/10 end
  return z==y
end
