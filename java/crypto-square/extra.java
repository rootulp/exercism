import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;

public class Crypto {

  private String input;

  public Crypto(String input) {
    this.input = input;
  }

  public String getNormalizedPlaintext() {
    return input.toLowerCase().replaceAll("[^\\w\\d]", "");
  }

  public int getSquareSize() {
    return (int) Math.ceil(Math.sqrt(getNormalizedPlaintext().length()));
  }

  public List<String> getPlaintextSegments() {
    String normalizedPlaintext = getNormalizedPlaintext();
    int squareSize = getSquareSize();

    List<String> results = new ArrayList<String>();

    int i;
    for (i = squareSize; i < normalizedPlaintext.length(); i += squareSize) {
      // Add chunks of squareSize num chars
      results.add(normalizedPlaintext.substring(i - squareSize, i));
    }
    // Add last chunk of characters regardless of size
    results.add(normalizedPlaintext.substring(i - squareSize));

    return results;
  }

  public String getCipherText() {
    return String.join("", getCipherSegments());
  }

  public String getNormalizedCipherText() {
    /* System.out.println(Arrays.toString(getCipherSegments().toArray())); */
    /* System.out.println(Arrays.toString(getPlaintextSegments().toArray())); */
    String result = "";
    String cipherText = getCipherText();
    int squareSize = getSquareSize();

    for (int i = 0, n = cipherText.length(); i < n; i++) {

      // add space after every GROUP_SIZE num of chars
      if (i != 0 && i % squareSize == 0) { result += " "; }

      result += cipherText.charAt(i);
    }

    return result;
  }

  private List<String> getCipherSegments() {
    List<String> plaintextSegments = getPlaintextSegments();
    int squareSize = getSquareSize();

    List<String> results = new ArrayList<String>();

    for (int i = 0; i < squareSize; i++) {
      String str = "";
      for (String segment : plaintextSegments) {
        if (i < segment.length()) { str += segment.charAt(i); }
      }
      results.add(str);
    }

    return results;
  }

}
