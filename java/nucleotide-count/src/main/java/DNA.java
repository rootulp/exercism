import java.util.Map;
import java.util.HashMap;

public final class DNA {

  private final String chain;
  private static final String VALID_NUCLEOTIDES = "ACGT";

  public DNA(String chain) {
    this.chain = chain;
  }

  public int count(char nucleotide) {
    if (invalid(nucleotide)) { throw new IllegalArgumentException(); }

    try {
      return nucleotideCounts().get(nucleotide);
    } catch (NullPointerException e) {
      return 0;
    }
  }

  private static boolean invalid(char nucleotide) {
    return VALID_NUCLEOTIDES.indexOf(nucleotide) == -1;
  }

  public Map<Character, Integer> nucleotideCounts() {
    Map<Character, Integer> counts = emptyCounts();
    for (char c : chain.toCharArray()) {
      counts.put(c, counts.get(c) + 1);
    }
    return counts;
  }

  private static Map<Character, Integer> emptyCounts() {
    Map<Character, Integer> emptyCounts = new HashMap<Character, Integer>();
    for (char c : VALID_NUCLEOTIDES.toCharArray()) {
      emptyCounts.put(c, 0);
    }
    return emptyCounts;
  }
}
