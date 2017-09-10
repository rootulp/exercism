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
    this.cipher = generateCipher(shift);
//    System.out.println(cipher);
  }

  private Map<Character, Character> generateCipher(int shift) {
    char[] lowercase_encrypted_alphabet = encrypt_alphabet(LOWERCASE_ALPHABET);
    char[] uppercase_encrypted_alphabet = encrypt_alphabet(UPPERCASE_ALPHABET);

    Map<Character, Character> cipher = new HashMap<>();

    for (int i = 0; i < ALPHABET_LENGTH; i++) {
      cipher.put(lowercase_encrypted_alphabet[i], LOWERCASE_ALPHABET[i]);
      cipher.put(uppercase_encrypted_alphabet[i], UPPERCASE_ALPHABET[i]);
    }

    return cipher;
  }

  private char[] encrypt_alphabet(char[] alphabet) {
    char[] encrypted = new char[ALPHABET_LENGTH];

    for (int i = 0; i < alphabet.length; i++) {
      encrypted[(i + shift) % alphabet.length] = alphabet[i];
    }

    return encrypted;
  }

  public String rotate(String text) {
    StringBuilder encrypted = new StringBuilder(text.length());
    for(char c: text.toCharArray()) {
      encrypted.append(encrypteChar(c));
    }
    return encrypted.toString();
  }

  private Character encrypteChar(Character c) {
    if (cipher.containsKey(c)) {
      return this.cipher.get(c);
    } else {
      return c;
    }
  }

}
