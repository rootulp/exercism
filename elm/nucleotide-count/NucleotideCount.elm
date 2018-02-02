module NucleotideCount exposing (..)

import Dict exposing (Dict)


type alias NucleotideCounts = {
  a : Int,
  t : Int,
  c : Int,
  g : Int
}

nucleotideCounts: String -> NucleotideCounts
nucleotideCounts strand =
  let occurencesInStrand = occurencesIn strand in
  { a = occurencesInStrand 'A'
  , t = occurencesInStrand 'T'
  , c = occurencesInStrand 'C'
  , g = occurencesInStrand 'G'
  }


occurencesIn: String -> Char -> Int
occurencesIn strand nucleotide =
  let nucleotideOccurences = occurences strand in
  nucleotideOccurences
    |> Dict.get nucleotide
    |> Maybe.withDefault 0


occurences: String -> Dict Char Int
occurences strand =
  strand
    |> String.toList
    |> List.foldl populateDict Dict.empty


populateDict : Char -> Dict Char Int -> Dict Char Int
populateDict nucleotide dict =
    Dict.update nucleotide incrementValue dict


incrementValue : Maybe Int -> Maybe Int
incrementValue maybe =
    Maybe.withDefault 0 maybe
        + 1
        |> Just


version : Int
version =
  2
