import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;

public class Anagram {

  private final String word;
  private final String sorted;

  public Anagram(String word) {
    char[] chars = word.toLowerCase().toCharArray();
    Arrays.sort(chars);

    this.word = word.toLowerCase();
    this.sorted = new String(chars);
  }

  public List<String> match(List<String> potentialWords) {

    List<String> anagrams = new ArrayList<String>();

    for (String word2 : potentialWords) {
      char[] chars = word2.toLowerCase().toCharArray();
      Arrays.sort(chars);
      String sorted2 = new String(chars);

      System.out.println("word = " + word);
      System.out.println("word2 = " + word2);
      System.out.println("sorted = " + sorted);
      System.out.println("sorted2 = " + sorted2);
      System.out.println("sorted equality = " + sorted2.equals( sorted));
      System.out.println("word equality = " + word2 != word);
      if (sorted2.toLowerCase().equals(sorted) && !word2.toLowerCase().equals(word)) { anagrams.add(word2); }
    }

    return anagrams;
  }

}
