import java.util.Arrays;
import java.util.stream.IntStream;

public class LargestSeriesProductCalculator {

  private long[] digits;

  public LargestSeriesProductCalculator(String series) {
    throwExceptionIfInvalidSeries(series);
    this.digits = parseDigits(series);
  }

  public long calculateLargestProductForSeriesLength(int seriesLength) {
    throwExceptionIfInvalidSeriesLength(seriesLength);
    return calculateLargestProduct(seriesLength);
  }

  private long calculateLargestProduct(int seriesLength) {
    long largestProduct = Long.MIN_VALUE;
    for (int i = 0; i <= digits.length - seriesLength; i++) {
      long product = calculateProductForSubseries(seriesLength, i);
      if (product > largestProduct) {
        largestProduct = product;
      }
    }
    return largestProduct;
  }

  private long calculateProductForSubseries(int seriesLength, int startingIndex) {
    long[] series = Arrays.copyOfRange(digits, startingIndex, startingIndex + seriesLength);
    return productOfSeries(series);
  }

  private static long productOfSeries (long[] series) {
    return Arrays.stream(series).reduce(1, (a, b) -> a * b);
  }

  private long[] parseDigits(String series) {
    // This is for https://github.com/exercism/problem-specifications/blob/master/exercises/largest-series-product/canonical-data.json#L89
    if (series != "") {
      return Arrays.stream(series.split("")).mapToLong(Long::parseLong).toArray();
    }
    return new long[]{};
  }

  private void throwExceptionIfInvalidSeries(String series) {
    if (series == null) {
      throw new IllegalArgumentException("String to search must be non-null.");
    }
    if (seriesContainsNonDigitCharacters(series)) {
      throw new IllegalArgumentException("String to search may only contains digits.");
    }
  }

  private boolean seriesContainsNonDigitCharacters(String series) {
    return !series.matches("\\d*");
  }

  private void throwExceptionIfInvalidSeriesLength(int seriesLength) {
    if (seriesLength > digits.length) {
      throw new IllegalArgumentException("Series length must be less than or equal to the length of the string to search.");
    }
    if (seriesLength < 0) {
      throw new IllegalArgumentException("Series length must be non-negative.");
    }
  }
}
