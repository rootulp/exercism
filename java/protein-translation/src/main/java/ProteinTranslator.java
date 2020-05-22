import java.util.List;
import java.util.Map;

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
		Map.entry("AUG", "Methionine")
	);

    List<String> translate(String rnaSequence) {
		return List.of(CODON_TO_PROTEIN.get(rnaSequence));
    }
}
