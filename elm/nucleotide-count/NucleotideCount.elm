module NucleotideCount exposing (..)

import Dict exposing (Dict)
-- import Record exposing (Record)
type alias NucleotideCounts = {
  a : Int,
  t : Int,
  c : Int,
  g : Int
}

nucleotideCounts: String -> NucleotideCounts
nucleotideCounts strand =
  let d = strand
    |> String.toList
    |> List.foldl populateDict Dict.empty in

  { a = Maybe.withDefault 0 (Dict.get 'A' d)
  , t = Maybe.withDefault 0 (Dict.get 'T' d)
  , c = Maybe.withDefault 0 (Dict.get 'C' d)
  , g = Maybe.withDefault 0 (Dict.get 'G' d)
  }

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
