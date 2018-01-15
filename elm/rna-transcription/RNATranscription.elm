module RNATranscription exposing (..)

toRNA: String -> Result Char String
toRNA strand =
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
    _ -> 'Z'
