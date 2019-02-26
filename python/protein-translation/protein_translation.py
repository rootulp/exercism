# Codon                 | Protein
# :---                  | :---
# AUG                   | Methionine
# UUU, UUC              | Phenylalanine
# UUA, UUG              | Leucine
# UCU, UCC, UCA, UCG    | Serine
# UAU, UAC              | Tyrosine
# UGU, UGC              | Cysteine
# UGG                   | Tryptophan
# UAA, UAG, UGA         | STOP

CODON_TO_PROTEIN = {
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan"
}

STOP_CODONS = set({"UAA", "UAG", "UGA"})
CODON_LENGTH = 3


def proteins(strand):
    polypeptide = []
    codons = split_into_codons(strand)

    for codon in codons:
        if (codon in STOP_CODONS):
            return polypeptide
        else:
            polypeptide.append(CODON_TO_PROTEIN[codon])

    return polypeptide


def split_into_codons(strand):
    return [strand[i:i + CODON_LENGTH] for i in range(0, len(strand),
                                                      CODON_LENGTH)]
