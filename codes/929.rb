# @param {String[]} emails
# @return {Integer}
def num_unique_emails(emails)
    emails.map do |email|
        pair = email.split("@")
        pair[0] = pair[0].delete(".")
        pair[0] = pair[0][0...pair[0].index("+")]
        pair.join("@")
    end.uniq.size
end
