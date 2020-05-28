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
    this.cipher = new HashMap<>();
    populateCipher();
  }

  public String rotate(String text) {
    StringBuilder encrypted = new StringBuilder(text.length());

    for (char c : text.toCharArray()) {
      encrypted.append(encryptChar(c));
    }

    return encrypted.toString();
  }

  private Character encryptChar(Character c) {
    return cipher.containsKey(c) ? this.cipher.get(c) : c;
  }

  private void populateCipher() {
    addAlphabetToCipher(LOWERCASE_ALPHABET);
    addAlphabetToCipher(UPPERCASE_ALPHABET);
  }

  private void addAlphabetToCipher(char[] alphabet) {
    for (int i = 0; i < alphabet.length; i++) {
      cipher.put(alphabet[i], alphabet[(i + shift) % alphabet.length]);
    }
  }
}
