class Pangram {

  sentence: string
  alphabet = new Set("ABCDEFGHIJKLMNOPQRSTUVWXYZ".split(""))

  constructor(sentence: string) {
    this.sentence = sentence
  }

  isPangram(): boolean {
    for (const character of this.sentence.toUpperCase()) {
      this.alphabet.delete(character)
    }
    return this.alphabet.size === 0
  }

}

export default Pangram