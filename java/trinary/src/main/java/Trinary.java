public class Trinary {

  private static int RADIX = 3;

  public static int toDecimal(String input) {
    if (!input.matches("[012]+")) {
      return 0;
    }
    return computeTrinary(input);
  }

  private static int computeTrinary(String input) {
    int result = 0;
    int pwr = 1;

    for (int i = input.length() - 1; i >= 0; i--, pwr *= RADIX) {
      int currDigit = Character.digit(input.charAt(i), RADIX);
      result += currDigit * pwr;
    }

    return result;
  }
}
