const COMPLEMENTS = new Map([
  ["G", "C"],
  ["C", "G"],
  ["T", "A"],
  ["A", "U"]
])

class Transcriptor {
  toRna(strand) {
    if (!this.validInput(strand)) {
      throw new Error("Invalid input DNA.");
    }
    return [...strand].map(nucleotide => COMPLEMENTS.get(nucleotide)).join("");
  }

  validInput(strand) {
    return strand.match(/^[GCTA]+$/);
  }
}

export default Transcriptor;