module Anagram exposing (..)

detect: String -> List String -> List String
detect word possibleAnagrams =
    List.filter (\ possibleAnagram -> List.sort (String.toList possibleAnagram) == List.sort (String.toList word)) possibleAnagrams
