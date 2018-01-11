module WordCount exposing (..)
import Dict exposing (Dict)

wordCount: String -> Dict String Int
wordCount sentence =
    String.words sentence
        |> List.map (\word -> (word, 1))
        |> List.foldl (\word dict -> Dict.update word Dict.get word dict)
