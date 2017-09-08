function Robot() {
  "use strict";

  var alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
  var digits = "0123456789";

  var randomString = function(chars, length) {
    var temp = "";
    for (var i = 0; i < length; i++) {
      temp += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return temp;
  };

  this.reset = function() {
    this.name = randomString(alpha, 2) + randomString(digits, 3);
  };

  this.reset();
}

module.exports = Robot;
