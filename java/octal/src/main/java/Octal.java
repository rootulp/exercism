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
    Integer decimal = 0;
    if (!input.matches("[0-7]+")) { return decimal; }

    for (int i = 0; i < input.length() - 1; i++) {
      int curr = Character.getNumericValue(input.charAt(i));
      int bla =  curr * (int) Math.pow(8, input.length() - i - 1);
      decimal += bla;
      System.out.println("curr " + curr + " bla " + bla);
    }

    return decimal;
  }
}
