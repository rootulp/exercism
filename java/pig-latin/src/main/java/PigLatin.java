import java.util.Set;
import java.util.HashSet;
import java.util.Arrays;

public class PigLatin {

  private static final Character[] VOWELS =
    new Character[] {'a', 'e', 'i', 'o', 'u'};
  private static final Set<Character> VOWELS_SET =
    new HashSet<Character>(Arrays.asList(VOWELS));

  public static String translate(String input) {
    if ((VOWELS_SET.contains(input.charAt(0))) ||
        (input.startsWith("xr")) ||
        (input.startsWith("yt"))) {
      return input + "ay";
    } else if (input.startsWith("squ") ||
               input.startsWith("sch") ||
               input.startsWith("thr")) {
      return input.substring(3) + input.substring(0, 3) + "ay";
    } else {
      return "bla";
    }
  }
}
