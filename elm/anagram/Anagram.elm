module Anagram exposing (..)

detect: String -> List String -> List String
detect word possibleAnagrams =
    List.filter (findAnagrams word) possibleAnagrams

findAnagrams: String -> (String -> Bool)
findAnagrams word =
    (\ possibleAnagram -> isAnagramCaseInsensitive word possibleAnagram)

isAnagramCaseInsensitive: String -> String -> Bool
isAnagramCaseInsensitive word possibleAnagram =
    isAnagram (String.toLower word) (String.toLower possibleAnagram)

isAnagram: String -> String -> Bool
isAnagram word possibleAnagram =
    word /= possibleAnagram &&
    sortedCharacters  word == sortedCharacters possibleAnagram

sortedCharacters: String -> List Char
sortedCharacters word =
    String.toList word
        |> List.sort
