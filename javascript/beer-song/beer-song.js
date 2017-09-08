var Beer = (function() {
  "use strict";

  function verse(num) {
    if (num === 0) {
      return (
        "No more bottles of beer on the wall, no more bottles of beer.\n" +
        "Go to the store and buy some more, " +
        "99 bottles of beer on the wall.\n"
      );
    } else if (num === 1) {
      return (
        "1 bottle of beer on the wall, 1 bottle of beer.\n" +
        "Take it down and pass it around, " +
        "no more bottles of beer on the wall.\n"
      );
    } else {
      return (
        num +
        " bottles of beer on the wall, " +
        num +
        " bottles of beer.\nTake one down and pass it around, " +
        (num - 1) +
        " bottles of beer on the wall.\n"
      );
    }
  }

  function sing(from, to) {
    to = to || 0;
    var verses = [];

    for (var i = from; i >= to; i--) {
      verses.push(this.verse(i));
    }

    return verses.join("\n");
  }

  return {
    verse: verse,
    sing: sing
  };
})();

module.exports = Beer;
