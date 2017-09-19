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
    return this.addCharCodes(char.charCodeAt(0), this.key.charCodeAt(index) - 97);
  }

  this.decodeChar = function(char, index) {
    return this.subtractCharCodes(char.charCodeAt(0), this.key.charCodeAt(index) - 97);
  }

  this.addCharCodes = function(charCode, encodeCharCode) {
    // console.log("charCode: " + charCode);
    // console.log("encodeCharCode: " + encodeCharCode);
    // console.log("wrapAround: " + this.wrapAround(charCode + encodeCharCode));
    return String.fromCharCode(this.wrapAround(charCode + encodeCharCode));
  }

  this.subtractCharCodes = function(charCode, encodeCharCode) {
    // console.log("charCode: " + charCode);
    // console.log("encodeCharCode: " + encodeCharCode);
    // console.log("wrapAround: " + this.wrapAround(charCode + encodeCharCode));
    return String.fromCharCode(this.wrapAround(charCode - encodeCharCode));
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
