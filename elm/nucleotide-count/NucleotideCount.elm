module NucleotideCount exposing (..)

-- import Record exposing (Record)
type alias NucleotideCounts = {
  a : Int,
  t : Int,
  c : Int,
  g : Int
}

nucleotideCounts: String -> NucleotideCounts
nucleotideCounts strand =
  { a = 0
  , t = 0
  , c = 0
  , g = 0
  }

version : Int
version =
  2
