import java.lang.Math;

public final class Octal {

  private String input;
  private int decimal;

  public Octal(String input) {
    this.input = input;
    decimal = calculateDecimal();
  }

  public int getDecimal() {
    return decimal;
  }

  private int calculateDecimal() {
    int decimal = 0;
    if (!input.matches("[0-7]+")) { return decimal; }

    for (int i = 0; i < input.length(); i++) {
      int curr = Character.getNumericValue(input.charAt(i));
      int mult =  (int) Math.pow(8, input.length() - 1 - i);
      decimal += (curr * mult);
    }
    return decimal;
  }
}
