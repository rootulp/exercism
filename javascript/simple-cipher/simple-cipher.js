function Cipher(key = generateKey()) {
  if (!key.match(/^[a-z]+$/)) {
    throw new Error("Bad key");
  }
  this.key = key;

  this.chars = function(str) {
    return str.split("");
  }

  this.encode = function(str) {
    return this.chars(str).map((char, index) => this.encodeChar(char, index)).join("");
  }

  this.decode = function(str) {
    return this.chars(str).map((char, index) => this.decodeChar(char, index)).join("");
  }

  this.encodeChar = function(char, index) {
    return this.add(char.charCodeAt(0), this.shiftDistance(index));
  }

  this.decodeChar = function(char, index) {
    return this.subtract(char.charCodeAt(0), this.shiftDistance(index));
  }

  this.add = function(charCode, shiftDistance) {
    return String.fromCharCode(this.wrapAround(charCode + shiftDistance));
  }

  this.subtract = function(charCode, shiftDistance) {
    return String.fromCharCode(this.wrapAround(charCode - shiftDistance));
  }

  this.shiftDistance = function(index) {
    return this.key.charCodeAt(index % this.key.length) - 97
  }

  // ASCII 97 = "a", 122 = "z"
  this.wrapAround = function(charCode) {
    if (charCode > 122) {
      return (charCode - 26);
    } else if (charCode < 97) {
      return (charCode + 26);
    } else {
      return charCode;
    }
  }
}

function generateKey() {
  return "aaaaaaaaaa"
}

module.exports = Cipher
