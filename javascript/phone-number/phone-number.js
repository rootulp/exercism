function PhoneNumber(raw_input) {
  'use strict';

  var invalid = '0000000000';

  function number() {
    number = strip_punctuation();

    if (number.length == 10) {
      return number;
    } else if (number.length == 11 && number[0] === '1') {
      return number.slice(1);
    } else {
      return invalid;
    }
  }

  function toString() {
    var areaCode = this.areaCode();
    var partOne = this.number().substr(3, 3);
    var partTwo = this.number().substr(6);
    return "(" + areaCode + ") " + partOne + "-" + partTwo;
  }

  function areaCode() {
    return number().substring(0, 3);
  }

  function strip_punctuation() {
    var re = /[\.\-\(\)\s]/g
    return raw_input.replace(re, '')
  }

  return {
    number: number,
    areaCode: areaCode,
    toString: toString
  }
}

module.exports = PhoneNumber
