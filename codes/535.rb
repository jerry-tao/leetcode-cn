# Encodes a URL to a shortened URL.
#
# @param {string} longUrl
# @return {string}
@cache = {}
def encode(longUrl)
    t = Time.now.to_s
    @cache[t] = longUrl
    t
end

# Decodes a shortened URL to its original URL.
#
# @param {string} shortUrl
# @return {string}
def decode(shortUrl)
    @cache[shortUrl]
end


# Your functions will be called as such:
# decode(encode(url))
