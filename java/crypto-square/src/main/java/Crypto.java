import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;

public class Crypto {

  private String normalizedPlaintext;
  private int squareSize;
  private int height;
  private char[][] matrix;
  private char[][] transposedMatrix;

  public Crypto(String input) {
    this.normalizedPlaintext = setNormalizedPlaintext(input);
    this.squareSize = setSquareSize(normalizedPlaintext);
    this.height = setHeight(normalizedPlaintext);

    this.matrix = buildMatrix();
    this.transposedMatrix = transposeMatrix();
  }

  public String getNormalizedPlaintext() {
    return normalizedPlaintext;
  }

  public int getSquareSize() {
    return squareSize;
  }

  public List<String> getPlaintextSegments() {
    return getSegments(matrix);
  }

  public String getCipherText() {
    StringBuilder sb = new StringBuilder();

    for (int i = 0; i < squareSize; i++) {
      for (int j = 0; j < squareSize; j++) {
        if (Character.isLetterOrDigit(transposedMatrix[i][j])) {
          sb.append(transposedMatrix[i][j]);
        }
      }
    }

    return sb.toString();
  }

  public String getNormalizedCipherText() {
    String cipherText = getCipherText();
    StringBuilder sb = new StringBuilder();

    for (int i = 0, n = cipherText.length(); i < n; i++) {
      if (i != 0 && i % height == 0) { sb.append(" "); }
      sb.append(cipherText.charAt(i));
    }

    return sb.toString();
  }

  private String setNormalizedPlaintext(String input) {
    return input.toLowerCase().replaceAll("[^\\w\\d]", "");
  }

  private int setSquareSize(String normaliezedPlaintext) {
    return (int) Math.ceil(Math.sqrt(normaliezedPlaintext.length()));
  }

  private int setHeight(String normaliezedPlaintext) {
    return (int) Math.floor(Math.sqrt(normaliezedPlaintext.length()));
  }

  private char[][] buildMatrix() {
    char[][] m = new char[squareSize][squareSize];

    for (int i = 0; i < normalizedPlaintext.length(); i++) {
      m[(int) Math.floor(i / squareSize)][i % squareSize] =
        normalizedPlaintext.charAt(i);
    }

    return m;
  }

  private char[][] transposeMatrix() {
    char[][] tp = buildMatrix();

    for (int i = 0; i < squareSize; i++) {
      for (int j = i + 1; j < squareSize; j++) {
        char temp = tp[i][j];
        tp[i][j] = tp[j][i];
        tp[j][i] = temp;
      }
    }

    return tp;
  }

  private List<String> getSegments(char[][] m) {
    List<String> segments = new ArrayList<String>();

    for (int i = 0; i < m.length; i++) {
      String segment = new String(m[i]).trim();
      if (!segment.isEmpty()) { segments.add(segment); }
    }

    return segments;
  }

}
