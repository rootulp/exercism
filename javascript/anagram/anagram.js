function anagram(subject) {
  var sortedSubject = sortWord(subject);

  function findMatches(words) {
    if (typeof words === "string") {
      words = Array.prototype.slice.call(arguments);
    }
    return words.filter(isAnagram);
  }

  function isAnagram(word) {
    return differentWord(subject, word) && sameSorted(subject, word);
  }

  function differentWord(subject, word) {
    return subject.toLowerCase() !== word.toLowerCase();
  }

  function sameSorted(subject, word) {
    var sortedWord = sortWord(word);
    return sortedSubject === sortedWord;
  }

  function sortWord(word) {
    return word
      .toLowerCase()
      .split("")
      .sort()
      .join("");
  }

  return {
    matches: findMatches
  };
}

module.exports = anagram;
