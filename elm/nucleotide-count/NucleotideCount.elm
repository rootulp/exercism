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
  { a = occurencesOf strand 'A'
  , t = occurencesOf strand 'T'
  , c = occurencesOf strand 'C'
  , g = occurencesOf strand 'G'
  }


occurencesOf: String -> Char -> Int
occurencesOf strand nucleotide =
  let occurences = nucleotideOccurences strand in
  occurences
    |> Dict.get nucleotide
    |> Maybe.withDefault 0


nucleotideOccurences: String -> Dict Char Int
nucleotideOccurences strand =
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
