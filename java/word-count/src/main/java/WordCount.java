import java.util.Map;
import java.util.HashMap;

public final class WordCount {

  public Map<String, Integer> Phrase(String words) {
    return countWords(words);
  }

  private static Map<String, Integer> countWords(String words) {
    Map<String, Integer> wordCounts = new HashMap<String, Integer>();

    /* Not a huge fan of the regex on the following line */
    String[] wordsArr = words.replaceAll("[^\\w\\s\\d]", "")
                              .toLowerCase()
                              .split("\\s+");

    for (String word : wordsArr) {
      if (wordCounts.containsKey(word)) {
        wordCounts.put(word, wordCounts.get(word) + 1);
      } else {
        wordCounts.put(word, 1);
      }
    }

    return wordCounts;
  }
}
