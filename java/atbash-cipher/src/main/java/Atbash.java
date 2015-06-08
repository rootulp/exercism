public class Atbash {

  private static final String PLAIN = "abcdefghijklmnopqrstuvwxyz";
  private static final String CIPHER = "zyxwvutsrqponmlkjihgfedcba";
  private static final int GROUP_SIZE = 5;

  public static String encode(String input) {

    // clean input
    String str = input.toLowerCase().replaceAll("[^\\w\\d]", "");
    String result = "";

    for (int i = 0, n = str.length(); i < n; i++) {

      // add space after every GROUP_SIZE num of chars
      if (i != 0 && i % GROUP_SIZE == 0) { result += " "; }

      result += cipherChar(str.charAt(i));
    }

    return result;
  }

  private static char cipherChar(char c) {
    int index = PLAIN.indexOf(c);

    // return original char if it doesn't exist in PLAIN (i.e nums)
    return index >= 0 ? CIPHER.charAt(index) : c;
  }

}
