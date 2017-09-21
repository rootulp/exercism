const DEFAULT_KEY = "aaaaaaaaaa"

class Cipher {

  constructor(key=DEFAULT_KEY) {
    if (!this.validKey(key)) {
      throw new Error("Bad key");
    }
    this.key = key
  }

  encode(plaintext) {
    return [...plaintext].map((char, index) => this.encodeChar(char, this.keyChar(index))).join("")
  }

  decode(plaintext) {
    return [...plaintext].map((char, index) => this.decodeChar(char, this.keyChar(index))).join("")
  }

  encodeChar(char, keyChar) {
    const charCode = char.charCodeAt(0)
    const keyCharCode = keyChar.charCodeAt(0)
    return String.fromCharCode(this.wrap(charCode + (keyCharCode - 97)));
  }

  decodeChar(char, keyChar) {
    const charCode = char.charCodeAt(0)
    const keyCharCode = keyChar.charCodeAt(0)
    return String.fromCharCode(this.wrap(charCode - (keyCharCode - 97)));
  }

  keyChar(index) {
    return this.key.charAt(index % this.key.length);
  }

  wrap(charCode) {
    if (charCode > 122) {
      return (charCode - 26);
    } else if (charCode < 97) {
      return (charCode + 26);
    } else {
      return charCode;
    }
  }

  validKey(key) {
    return key.match(/^[a-z]+$/);
  }
}

export default Cipher;
