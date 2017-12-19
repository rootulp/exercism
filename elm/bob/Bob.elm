module Bob exposing (..)

hey: String -> String
hey input =
    if shouting input then
        "Whoa, chill out!"
    else if question input then
         "Sure."
    else
        "Whatever."

shouting: String -> Bool
shouting input =
    String.toUpper input == input &&
    String.toLower input /= input


question: String -> Bool
question input =
    String.endsWith "?" input
