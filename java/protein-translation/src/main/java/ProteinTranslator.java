import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Set;

class ProteinTranslator {
  private static final Integer CODON_LENGTH = 3;
  private static final Map<String, String> CODON_TO_PROTEIN =
      Map.ofEntries(
          Map.entry("AUG", "Methionine"),
          Map.entry("UUU", "Phenylalanine"),
          Map.entry("UUC", "Phenylalanine"),
          Map.entry("UUA", "Leucine"),
          Map.entry("UUG", "Leucine"),
          Map.entry("UCU", "Serine"),
          Map.entry("UCC", "Serine"),
          Map.entry("UCA", "Serine"),
          Map.entry("UCG", "Serine"),
          Map.entry("UAU", "Tyrosine"),
          Map.entry("UAC", "Tyrosine"),
          Map.entry("UGU", "Cysteine"),
          Map.entry("UGC", "Cysteine"),
          Map.entry("UGG", "Tryptophan"));
  private static final Set<String> STOP_CODONS = Set.of("UAA", "UAG", "UGA");

  public List<String> translate(final String rnaSequence) {
    final List<String> codons = splitIntoCodons(rnaSequence);
    List<String> proteins = new ArrayList<>();
    for (String codon : codons) {
      if (STOP_CODONS.contains(codon)) {
        return proteins;
      }
      proteins.add(translateCodon(codon));
    }
    ;
    return proteins;
  }

  private static List<String> splitIntoCodons(final String rnaSequence) {
    final List<String> codons = new ArrayList<>();
    for (int i = 0; i < rnaSequence.length(); i += CODON_LENGTH) {
      codons.add(rnaSequence.substring(i, Math.min(rnaSequence.length(), i + CODON_LENGTH)));
    }
    return codons;
  }

  private static String translateCodon(final String codon) {
    return CODON_TO_PROTEIN.get(codon);
  }
}
