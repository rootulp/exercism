class Squares {

  public squareOfSums: number
  public sumOfSquares: number
  public difference: number

  constructor(num: number) {
    this.squareOfSums = Squares.squareOfSums(num)
    this.sumOfSquares = Squares.sumOfSquares(num)
    this.difference = this.squareOfSums - this.sumOfSquares
  }

  private static squareOfSums(num: number): number {
    return Squares.square(Squares.sum(Squares.range(num)))
  }

  private static sumOfSquares(num: number): number {
    return Squares.sum(Squares.range(num).map( (x) => Squares.square(x)))
  }

  private static range(num: number): number[] {
    return [...Array(num + 1).keys()]
  }

  private static sum(num: number[]): number {
    return num.reduce((a, b) => a + b, 0)
  }

  private static square(num: number): number {
    return num ** 2
  }

}

export default Squares