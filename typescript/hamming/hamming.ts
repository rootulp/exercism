class Hamming {

  public compute(strand1: string, strand2: string) {
    if (this.unequalLength(strand1, strand2)) {
      throw 'DNA strands must be of equal length.'
    }
    return this.distanceBetweenStrands(strand1, strand2)
  }

  private distanceBetweenStrands(strand1: string, strand2: string) {
    let distance = 0
    for (let i = 0; i < strand1.length; i++) {
      distance += this.distanceBetweenNucleotides(strand1.charAt(i), strand2.charAt(i))
    }
    return distance
  }

  private distanceBetweenNucleotides(nucleotide1: string, nucleotide2: string) {
    return nucleotide1 === nucleotide2 ? 0 : 1
  }

  private unequalLength(strand1: string, strand2: string) {
    return strand1.length !== strand2.length
  }

}

export default Hamming