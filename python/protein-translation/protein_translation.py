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


def proteins(strand):
    if (strand in STOP_CODONS):
        return []
    return [CODON_TO_PROTEIN[strand]]
