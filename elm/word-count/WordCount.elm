module WordCount exposing (..)
import Dict exposing (Dict)
import Char

wordCount: String -> Dict String Int
wordCount sentence =
    sentence
        |> removePunctuation
        |> String.words
        |> List.foldl populateDict Dict.empty

removePunctuation: String -> String
removePunctuation sentence =
    sentence
        |> String.toLower
        |> String.filter isLetterDigitOrWhitespace

isLetterDigitOrWhitespace: Char -> Bool
isLetterDigitOrWhitespace char =
    Char.isLower char ||
    Char.isDigit char ||
    char == ' '

populateDict: comparable -> Dict comparable Int -> Dict comparable Int
populateDict word dict =
    Dict.update word incrementValue dict

incrementValue : Maybe Int -> Maybe Int
incrementValue maybe =
    (Maybe.withDefault 0 maybe) + 1
        |> Just
