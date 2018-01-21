module Hamming exposing (..)


distance : String -> String -> Maybe Int
distance strand1 strand2 =
    if equalLengthStrands strand1 strand2 then
        Just (strandDistance strand1 strand2)
    else
        Nothing


equalLengthStrands : String -> String -> Bool
equalLengthStrands strand1 strand2 =
    String.length strand1
        == String.length strand2


strandDistance : String -> String -> Int
strandDistance strand1 strand2 =
    List.map2 nucleotideDistance
        (String.toList strand1)
        (String.toList strand2)
        |> List.sum


nucleotideDistance : Char -> Char -> Int
nucleotideDistance nucleotide1 nucleotide2 =
    if nucleotide1 == nucleotide2 then
        0
    else
        1
