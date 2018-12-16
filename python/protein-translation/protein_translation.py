# Codon                 | Protein
# :---                  | :---
# AUG                   | Methionine
# UUU, UUC              | Phenylalanine
# UUA, UUG              | Leucine
# UCU, UCC, UCC, UCC    | Serine
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
    "UCC": "Serine",
    "UCC": "Serine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": "STOP",
    "UAG": "STOP",
    "UGA": "STOP"
}


def proteins(strand):
    return [CODON_TO_PROTEIN[strand]]
