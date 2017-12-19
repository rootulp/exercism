module Bob exposing (..)

hey: String -> String
hey input =
    if question input then
         "Sure."
    else
        "Whatever."

question: String -> Bool
question input =
    String.endsWith "?" input
