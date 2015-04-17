function Gigasecond(dob) {
  'use strict';

  this.date = function() {
    var gigaDate = new Date(1e12 + dob.getTime());
    return new Date(gigaDate.toDateString());
  }

}

module.exports = Gigasecond
