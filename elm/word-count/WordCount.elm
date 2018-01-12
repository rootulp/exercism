module WordCount exposing (..)
import Dict exposing (Dict)

wordCount: String -> Dict String Int
wordCount sentence =
    String.words sentence
        |> List.foldl populateDict Dict.empty

populateDict: comparable -> Dict comparable Int -> Dict comparable Int
populateDict word dict =
    Dict.update word incrementValue dict

incrementValue : Maybe Int -> Maybe Int
incrementValue maybe =
    (Maybe.withDefault 0 maybe) + 1
        |> Just

