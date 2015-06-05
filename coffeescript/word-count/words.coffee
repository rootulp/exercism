class Words
  constructor: (input) ->
    counts = {}
    words = input.toLowerCase().match(/\b[\w\d]+\b/g)

    for word in words
      counts[word] = if counts[word] then counts[word] + 1 else 1

    @count = counts
  module.exports = @
