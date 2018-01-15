module RNATranscription exposing (..)

toRNA: String -> Result Char String
toRNA strand =
    if valid strand then
        transcribe strand
    else
        invalid strand

valid: String -> Bool
valid strand =
    strand
        |> String.toList
        |> List.all isDNANucleotide

-- It's not clear to me why the error handling tests
-- ask for the first Char in the strand
-- https://github.com/exercism/elm/issues/174
invalid: String -> Result Char String
invalid strand =
    strand
        |> String.toList
        |> List.head
        |> Maybe.withDefault 'X' -- Ideally remove this default
        |> Err

transcribe: String -> Result Char String
transcribe strand =
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
    List.member nucleotide ['G', 'C', 'T', 'A'] -- Ideally remove this duplication
