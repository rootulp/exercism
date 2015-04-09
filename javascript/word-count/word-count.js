function words(phrase) {
  'use strict';

  var counts = {};
  var words = phrase.split(/\s/);

  words.forEach(function(word) {
    if (counts[word]) {
      counts[word] += 1;
    } else {
      counts[word] = 1;
    }
  });

  return counts;
};

module.exports = words;
