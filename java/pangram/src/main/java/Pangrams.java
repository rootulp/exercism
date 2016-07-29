import java.util.Set;
import java.util.HashSet;

public class Pangrams {

  private static Set<Character> alphabet() {
    Set<Character> alphabet = new HashSet();
    for(char letter = 'a'; letter <= 'z'; letter++) {
      alphabet.add(letter);
    }
    return alphabet;
  }

  public static boolean isPangram(String candidate) {
    Set <Character> alphabet = alphabet();
    for (char c : candidate.toLowerCase().toCharArray()) {
      alphabet.remove(c);
    }
    return alphabet.isEmpty();
  }

}