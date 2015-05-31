import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;

public class Anagram {

  private final String word;
  private final String sorted;

  public Anagram(String word) {
    this.word = word;
    this.sorted = sort(word);
  }

  public List<String> match(List<String> potentialAnagrams) {

    List<String> anagrams = new ArrayList<String>();

    for (String potentialAnagram : potentialAnagrams) {
      String sortedPotentialAnagram = sort(potentialAnagram);

      if (sorted.equals(sortedPotentialAnagram) &&
          !word.equals(potentialAnagram.toLowerCase())) {
        anagrams.add(potentialAnagram);
      }
    }

    return anagrams;
  }

  private String sort(String word) {
    char[] chars = word.toLowerCase().toCharArray();
    Arrays.sort(chars);

    return new String(chars);
  }
}
