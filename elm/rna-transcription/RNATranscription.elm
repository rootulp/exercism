module RNATranscription exposing (..)

toRNA: String -> Result Char String
toRNA strand =
    if invalidInput strand then
        complementStrand strand
    else
        invalidStrand strand

invalidInput: String -> Bool
invalidInput strand =
    strand
        |> String.toList
        |> List.all isDNANucleotide

-- It's not clear to me why the error handling tests
-- ask for the first Char in the strand
-- https://github.com/exercism/elm/issues/174
invalidStrand: String -> Result Char String
invalidStrand strand =
    strand
        |> String.toList
        |> List.head
        |> Maybe.withDefault 'X'
        |> Err

complementStrand: String -> Result Char String
complementStrand strand =
    strand
        |> String.map complement
        |> Ok

complement: Char -> Char
complement nucleotide =
    case nucleotide of
    'G' -> 'C'
    'C' -> 'G'
    'T' -> 'A'
    'A' -> 'U'
    _ -> 'Z' -- Ideally strongly type Nucleotide and remove this case

isDNANucleotide: Char -> Bool
isDNANucleotide nucleotide =
    List.member nucleotide ['G', 'C', 'T', 'A']
