module.exports = class Anagram
  constructor: (w) ->
    @detector_word = w.toLowerCase()
    @detector_sorted = @detector_word.split('').sort().join()

  match: (words) ->
    words = words.map (word) -> word.toLowerCase()
    words.filter (word) =>
      return false if word is @detector_word
      word.split('').sort().join() is @detector_sorted
