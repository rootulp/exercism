module Pangram exposing (..)

import Set exposing (..)

isPangram: String -> Bool
isPangram sentence =
    let
        alphabet: Set Char
        alphabet = Set.fromList (String.toList "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    in
        Set.isEmpty (Set.diff alphabet (Set.fromList (String.toList (String.toUpper sentence))))
