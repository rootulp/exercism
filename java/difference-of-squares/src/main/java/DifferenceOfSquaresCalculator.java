import java.util.stream.IntStream;

public final class DifferenceOfSquaresCalculator {

  public int computeSquareOfSumTo(int num) {
    return (int) Math.pow(IntStream.rangeClosed(0, num).sum(), 2);
  }

  public int computeSumOfSquaresTo(int num) {
    return IntStream.rangeClosed(0, num).map(x -> (int) Math.pow(x, 2)).sum();
  }

  public int computeDifferenceOfSquares(int num) {
    return Math.abs(computeSquareOfSumTo(num) - computeSumOfSquaresTo(num));
  }
}
