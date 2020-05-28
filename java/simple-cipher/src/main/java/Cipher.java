import java.util.Random;

public final class Cipher {

  private String key;
  private static final String ALPHABET = "abcdefghijklmnopqrstuvwxyz";

  public Cipher(String key) {
    setKey(key);
  }

  public Cipher() {
    setKey(randomKey());
  }

  public String getKey() {
    return key;
  }

  private void setKey(String key) {
    if (!key.matches("\\p{javaLowerCase}+")) {
      throw new IllegalArgumentException();
    }

    this.key = key;
  }

  public String encode(String input) {
    StringBuilder output = new StringBuilder(input.length());

    for (int i = 0; i < Math.min(input.length(), key.length()); i++) {
      output.append(encodeChar(input.charAt(i), i));
    }

    return output.toString();
  }

  public String decode(String input) {
    StringBuilder output = new StringBuilder(input.length());

    for (int i = 0; i < input.length(); i++) {
      output.append(decodeChar(input.charAt(i), i));
    }

    return output.toString();
  }

  private char encodeChar(Character c, Integer i) {
    int alphaI = ALPHABET.indexOf(c) + ALPHABET.indexOf(key.charAt(i));
    return ALPHABET.charAt(wrapAlphaI(alphaI));
  }

  private char decodeChar(Character c, Integer i) {
    int alphaI = ALPHABET.indexOf(c) - ALPHABET.indexOf(key.charAt(i));
    return ALPHABET.charAt(wrapAlphaI(alphaI));
  }

  private Integer wrapAlphaI(Integer alphaI) {

    while (alphaI < 0) {
      alphaI += ALPHABET.length();
    }

    while (alphaI >= ALPHABET.length()) {
      alphaI -= ALPHABET.length();
    }

    return alphaI;
  }

  private String randomKey() {
    StringBuilder key = new StringBuilder(100);

    for (int i = 0; i < 100; i++) {
      key.append(randomChar());
    }

    return key.toString();
  }

  private char randomChar() {
    Random r = new Random();
    return (char) (r.nextInt(26) + 'a');
  }
}
