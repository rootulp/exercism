import org.apache.commons.lang3.StringUtils;

public class Matrix {

  int[][] matrix;

  public Matrix(String matrixAsString) {
    matrix = populateMatrix(matrixAsString);
  }

  public int[] getRow(int rowIndex) {
    return matrix[rowIndex];
  }

  public int[] getColumn(int colIndex) {
    int[][] transposedMatrix = transposeMatrix(matrix);
    return transposedMatrix[colIndex];
  }

  public int getRowsCount() {
    return matrix.length;
  }

  public int getColumnsCount() {
    return matrix[0].length;
  }

  private int[][] populateMatrix(String matrixAsString) {
    int rowCount = rowCount(matrixAsString);
    int columnCount = columnCount(matrixAsString);
    int[][] matrix = new int[rowCount][columnCount];

    String[] rows = matrixAsString.split("\n");
    for (int rowIndex = 0; rowIndex < rows.length; rowIndex += 1) {
      String row = rows[rowIndex];
      String[] columns = row.split(" ");
      for (int columnIndex = 0; columnIndex < columns.length; columnIndex += 1) {
        matrix[rowIndex][columnIndex] = Integer.parseInt(columns[columnIndex]);
      }
    }

    return matrix;
  }

  private int[][] transposeMatrix(int[][] matrix) {
    int[][] transposedMatrix = new int[getColumnsCount()][getRowsCount()];

    for (int rowIndex = 0; rowIndex < getRowsCount(); rowIndex += 1) {
      for (int columnIndex = 0; columnIndex < getColumnsCount(); columnIndex += 1) {
        transposedMatrix[columnIndex][rowIndex] = matrix[rowIndex][columnIndex];
      }
    }

    return transposedMatrix;
  }

  private int rowCount(String matrixAsString) {
    return StringUtils.countMatches(matrixAsString, "\n") + 1;
  }

  private int columnCount(String matrixAsString) {
    return StringUtils.countMatches(firstRow(matrixAsString), " ") + 1;
  }

  private String firstRow(String matrixAsString) {
    return matrixAsString.split( "\n")[0];
  }

  private void print(int[][] matrix) {
    for (int rowIndex = 0; rowIndex < getRowsCount(); rowIndex += 1) {
      for (int columnIndex = 0; columnIndex < getColumnsCount(); columnIndex += 1) {
        System.out.print(matrix[rowIndex][columnIndex] + " ");
      }
      System.out.println();
    }
    System.out.println("-----");
  }

}
