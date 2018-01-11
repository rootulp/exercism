module WordCount exposing (..)
import Dict exposing (Dict)

wordCount: String -> Dict String Int
wordCount sentence =
    String.words sentence
        |> List.map (\value -> (value, 1))
        |> Dict.fromList
