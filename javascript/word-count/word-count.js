var words = function(phrase) {
  var counts = {};
  phrase.split(/\s/).forEach(function(word) {
    if (counts[word]) {
      counts[word]++;
    } else {
      counts[word] = 1;
    }
  });
  return counts;
};

module.exports = words;
