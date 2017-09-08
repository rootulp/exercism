class Squares {

  public squareOfSums: number;
  public sumOfSquares: number;
  public difference: number;

  constructor(number: number) {
    this.squareOfSums = Squares.squareOfSums(number)
    this.sumOfSquares = Squares.sumOfSquares(number)
    this.difference = this.squareOfSums - this.sumOfSquares
  }

  private static squareOfSums(number: number): number {
    return Squares.square(Squares.sum(Squares.range(number)));
  }

  private static sumOfSquares(number: number): number {
    return Squares.sum(Squares.range(number).map( x => Squares.square(x)));
  }

  private static range(number: number): Array<number> {
    return [...Array(number + 1).keys()]
  }

  private static sum(numbers: Array<number>): number {
    return numbers.reduce((a, b) => a + b, 0)
  }

  private static square(number: number): number {
    return number ** 2;
  }

}

export default Squares;