class Matrix {

  private matrix: number[][];
  constructor(input: string) {
    this.matrix = input.split('\n').map(row => row.split(' ')).map(row => row.map(val => Number.parseInt(val)))
  }

  public get rows(): number[][] {
    return this.matrix;
  }

  public get columns(): number[][] {
    // TODO transpose
    return this.matrix;
  }
}

export default Matrix
