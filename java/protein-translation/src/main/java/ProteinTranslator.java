import java.util.List;
import java.util.Map;
import java.util.Set;

class ProteinTranslator {
	// Codon                 | Protein
	// :---                  | :---
	// AUG                   | Methionine
	// UUU, UUC              | Phenylalanine
	// UUA, UUG              | Leucine
	// UCU, UCC, UCA, UCG    | Serine
	// UAU, UAC              | Tyrosine
	// UGU, UGC              | Cysteine
	// UGG                   | Tryptophan
	// UAA, UAG, UGA         | STOP
	private static final Map<String, String> CODON_TO_PROTEIN = Map.ofEntries(
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
		Map.entry("UGG", "Tryptophan")
	);
	private static final Set<String> STOP_CODONS = Set.of("UAA", "UAG", "UGA");

    List<String> translate(String rnaSequence) {
		return List.of(CODON_TO_PROTEIN.get(rnaSequence));
    }
}
