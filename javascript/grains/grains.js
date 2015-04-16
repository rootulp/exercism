function Grains() {
  'use strict';

  this.square = function(num) {
    return Math.pow(2, num - 1);
  }

  this.total = function() {
    var result = 0;
    for (var i = 0; i <= 64; i++) {
      result += this.square(i);
    }
    return result;
  }

}

module.exports = Grains
