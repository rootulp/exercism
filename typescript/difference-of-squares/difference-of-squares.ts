class Squares {

  private number: number;
  public squareOfSums: number;
  public sumOfSquares: number;
  public difference: number;

  constructor(number: number) {
    this.number = number;
    this.squareOfSums = Squares.squareOfSums(number)
    this.sumOfSquares = Squares.sumOfSquares(number)
    this.difference = this.squareOfSums - this.sumOfSquares
  }

  private static squareOfSums(number: number): number {
    const range = Squares.range(number);
    const sum = range.reduce((a, b) => a + b, 0);
    const square = sum ** 2;
    return square;
  }

  private static sumOfSquares(number: number): number {
    const range = Squares.range(number);
    const squares = range.map( x => x ** 2);
    const sum = squares.reduce((a, b) => a + b, 0);
    return sum;
  }

  private static range(number: number): Array<number> {
    return [...Array(number + 1).keys()]
  }

}

export default Squares;