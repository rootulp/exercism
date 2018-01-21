module RunLengthEncoding exposing (..)

import Char
import Maybe
import Tuple


encode : String -> String
encode data =
    data
        |> String.toList
        |> List.foldr splitRuns []
        |> List.map encodeRun
        |> String.join ""


decode : String -> String
decode data =
    data
        |> String.toList
        |> List.foldr splitRuns []
        |> List.map decodeRun
        |> String.join ""


splitRuns : Char -> List String -> List String
splitRuns value runs =
    if runBoundary value runs then
        startNewRun value runs
    else
        addValueToExistingRun value runs


runBoundary : Char -> List String -> Bool
runBoundary value runs =
    not
        (Char.isDigit value
            || -- during decode
               valueInPreviousRun value runs
        )



-- during encode


startNewRun : Char -> List String -> List String
startNewRun value runs =
    String.fromChar value :: runs


valueInPreviousRun : Char -> List String -> Bool
valueInPreviousRun value runs =
    previousRun runs
        |> String.contains (String.fromChar value)


addValueToExistingRun : Char -> List String -> List String
addValueToExistingRun value runs =
    appendValueToPreviousRun value runs :: restOfRuns runs


appendValueToPreviousRun : Char -> List String -> String
appendValueToPreviousRun value runs =
    runs
        |> previousRun
        |> String.cons value


encodeRun : String -> String
encodeRun run =
    run
        |> String.left 1
        |> String.append (calculateRunLength run)


decodeRun : String -> String
decodeRun run =
    run
        |> splitIntoLengthAndValue
        |> expand


splitIntoLengthAndValue : String -> ( Int, String )
splitIntoLengthAndValue run =
    run
        |> String.toList
        |> List.partition Char.isDigit
        |> Tuple.mapFirst convertRunLength
        |> Tuple.mapSecond String.fromList


expand : ( Int, String ) -> String
expand ( length, value ) =
    String.repeat length value


convertRunLength : List Char -> Int
convertRunLength runLength =
    runLength
        |> String.fromList
        |> String.toInt
        |> Result.withDefault 1


calculateRunLength : String -> String
calculateRunLength run =
    let
        runLength =
            String.length run
    in
    if runLength == 1 then
        ""
    else
        toString runLength


previousRun : List String -> String
previousRun runs =
    runs
        |> List.head
        |> Maybe.withDefault ""


restOfRuns : List String -> List String
restOfRuns runs =
    runs
        |> List.tail
        |> Maybe.withDefault []


version : Int
version =
    2
