module Hamming exposing (..)

distance: String -> String -> Maybe Int
distance strand1 strand2 =
    List.map2 nucleotideDistance
        (String.toList strand1)
        (String.toList strand2)
        |> List.sum
        |> Just

nucleotideDistance: Char -> Char -> Int
nucleotideDistance nucleotide1 nucleotide2 =
    if nucleotide1 == nucleotide2 then
        0
    else
        1
