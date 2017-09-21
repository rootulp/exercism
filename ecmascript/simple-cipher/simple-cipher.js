
class Cipher {

  constructor(key=this.defaultKey()) {
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
    return String.fromCharCode(charCode + (keyCharCode - 97));
  }

  decodeChar(char, keyChar) {
    const charCode = char.charCodeAt(0)
    const keyCharCode = keyChar.charCodeAt(0)
    return String.fromCharCode(charCode - (keyCharCode - 97));
  }

  keyChar(index) {
    return this.key.charAt(index % this.key.length);
  }

  defaultKey() {
    return "aaaaaaaaaa"
  }
}

export default Cipher;