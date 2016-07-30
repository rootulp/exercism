import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

public class PigLatin {

  private static final Character[] VOWELS =
    new Character[] {'a', 'e', 'i', 'o', 'u'};
  private static final Set<Character> VOWELS_SET =
    new HashSet<Character>(Arrays.asList(VOWELS));

  public static String translate(String phrase) {
    String[] words = phrase.split(" ");
    for (int i = 0; i < words.length; i++) {
      words[i] = translateWord(words[i]);
    }
    return String.join(" ", words);
  }

  private static String translateWord(String word) {
    if (singleVowel(word)){
      return word + "ay";
    } else if (tripleConsonant(word)) {
      return word.substring(3) + word.substring(0, 3) + "ay";
    } else if (doubleConsonant(word)) {
      return word.substring(2) + word.substring(0, 2) + "ay";
    } else {
      return word.substring(1) + word.substring(0, 1) + "ay";
    }
  }

  private static boolean singleVowel(String word) {
    return ((VOWELS_SET.contains(word.charAt(0))) ||
            (word.startsWith("xr")) ||
            (word.startsWith("yt")));
  }

  private static boolean tripleConsonant(String word) {
    return (word.startsWith("squ") ||
            word.startsWith("sch") ||
            word.startsWith("thr"));
  }

  private static boolean doubleConsonant(String word) {
    return (word.startsWith("ch") ||
            word.startsWith("qu") ||
            word.startsWith("th"));
  }
}
