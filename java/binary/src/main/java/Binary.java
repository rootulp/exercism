import java.lang.Math;

public class Binary {

  private final String input;

  public Binary(String input) {
    this.input = input;
  }

  public int getDecimal() {

    // Remove non 1's and 0's
    String cleanInput = input.replaceAll("[^01]","");

    // Return 0 if invalid string
    if (!input.equals(cleanInput)) { return 0; }

    // Reverse cleanInput
    String reverse = new StringBuffer(cleanInput).reverse().toString();

    // Increment total based on which indexes have 1's
    int total = 0;
    for (int i = 0, n = reverse.length(); i < n; i++) {
      if (reverse.charAt(i) == '1') { total += Math.pow(2, i); }
    }

    // Return total
    return total;
  }

}
