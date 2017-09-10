import java.util.HashMap;
import java.util.Map;

public class RotationalCipher {

  private int shift;
  private Map<Character, Character> cipher;

  private static int ALPHABET_LENGTH = 26;
  private static char[] LOWERCASE_ALPHABET = "abcdefghijklmnopqrstuvwxyz".toCharArray();
  private static char[] UPPERCASE_ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".toCharArray();

  public RotationalCipher(int shift) {
    this.shift = shift;
    this.cipher = generateCipher();
  }

  public String rotate(String text) {
    StringBuilder encrypted = new StringBuilder(text.length());

    for(char c: text.toCharArray()) {
      encrypted.append(encryptChar(c));
    }

    return encrypted.toString();
  }

  private Map<Character, Character> generateCipher() {
    char[] lowercaseRotatedAlphabet = rotateAlphabet(LOWERCASE_ALPHABET);
    char[] uppercaseRotatedAlphabet = rotateAlphabet(UPPERCASE_ALPHABET);

    Map<Character, Character> cipher = new HashMap<>();

    for (int i = 0; i < ALPHABET_LENGTH; i++) {
      cipher.put(lowercaseRotatedAlphabet[i], LOWERCASE_ALPHABET[i]);
      cipher.put(uppercaseRotatedAlphabet[i], UPPERCASE_ALPHABET[i]);
    }

    return cipher;
  }

  private char[] rotateAlphabet(char[] alphabet) {
    char[] rotated = new char[alphabet.length];

    for (int i = 0; i < alphabet.length; i++) {
      rotated[(i + shift) % alphabet.length] = alphabet[i];
    }

    return rotated;
  }

  private Character encryptChar(Character c) {
    return cipher.containsKey(c) ? this.cipher.get(c) : c;
  }

}
