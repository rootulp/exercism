module Pangram exposing (..)

import Set exposing (..)


isPangram : String -> Bool
isPangram sentence =
    String.toUpper sentence
        |> String.toList
        |> Set.fromList
        |> Set.diff alphabet
        |> Set.isEmpty


alphabet : Set Char
alphabet =
    String.toList "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        |> Set.fromList
