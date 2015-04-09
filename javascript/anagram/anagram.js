function anagram(subject) {

  var sortedSubject = sortWord(subject)

  function findMatches(words) {
    if (typeof words === 'string') {
      words = Array.prototype.slice.call(arguments)
    }
    return words.filter(isAnagram)
  }

  function sortWord(word) {
    return word.toLowerCase().split('').sort().join('');
  }

  function isAnagram(word) {
    var sortedWord = sortWord(word)

    if (subject.toLowerCase() !== word.toLowerCase()) {
      if (sortedSubject === sortedWord) {
        return true
      }
    }
    return false

  }

  return {
    matches: findMatches
  }
}

module.exports = anagram
