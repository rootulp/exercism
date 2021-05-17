class Matrix {

  private matrix: number[][];
  private transposed: number[][];
  constructor(input: string) {
    this.matrix = input.split('\n').map(row => row.split(' ')).map(row => row.map(val => Number.parseInt(val)))
    this.transposed = this.transpose(this.clone(this.matrix))
  }

  public get rows(): number[][] {
    return this.matrix;
  }

  public get columns(): number[][] {
    return this.transposed;
  }

  private transpose(matrix: number[][]): number[][] {
    return matrix[0].map((_, colIndex) => matrix.map(row => row[colIndex]));
  }
  private clone(matrix: number[][]): number[][] {
    let clone: number[][] = [];
    matrix.forEach(row => {
      let clonedRow: number[] = [];
      row.forEach(val => {
        clonedRow.push(val)
      })
      clone.push(clonedRow);
    });
    return clone;
  }
}

export default Matrix
