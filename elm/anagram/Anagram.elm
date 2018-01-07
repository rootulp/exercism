module Anagram exposing (..)

detect: String -> List String -> List String
detect word possibleAnagrams =
    List.filter (\ possibleAnagram ->  isAnagram (String.toLower possibleAnagram) (String.toLower word)) possibleAnagrams

isAnagram: String -> String -> Bool
isAnagram word possibleAnagram =
     word /= possibleAnagram &&
    sortedCharacters  word == sortedCharacters possibleAnagram

sortedCharacters: String -> List Char
sortedCharacters word =
    String.toList word
        |> List.sort
