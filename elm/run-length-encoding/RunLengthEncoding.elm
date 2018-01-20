module RunLengthEncoding exposing (..)
import Maybe
import Char
import Tuple

encode: String -> String
encode data =
  data
    |> String.toList
    |> List.foldr splitRuns []
    |> List.map encodeRun
    |> String.join ""

decode: String -> String
decode data =
  data
    |> String.toList
    |> List.foldr splitEncodedRuns []
    |> List.map decodeRun
    |> String.join ""

splitRuns: (Char -> List String -> List String)
splitRuns char runs =
    if existingRun char runs then
      handleExistingRun char runs
    else
      newRun char runs

splitEncodedRuns: (Char -> List String -> List String)
splitEncodedRuns char runs =
  let nextChar = String.fromChar char in
    if Char.isDigit char then
      handleExistingRun char runs
    else
      newRun char runs

newRun: Char -> List String -> List String
newRun char runs =
  let nextChar = String.fromChar char in
    nextChar :: runs

existingRun: Char -> List String -> Bool
existingRun char runs =
  String.startsWith (String.fromChar char) (previousRun runs)

handleExistingRun: Char -> List String -> List String
handleExistingRun char runs =
  runs
    |> restOfRuns
    |> (::) (appendCharToExistingRun char runs)

appendCharToExistingRun: Char -> List String -> String
appendCharToExistingRun char runs =
  runs
    |> previousRun
    |> String.cons char

encodeRun: String -> String
encodeRun run =
    run
      |> String.left 1
      |> String.append (calculateRunLength run)

decodeRun: String -> String
decodeRun run =
  run
    |> decodeLengthAndChar
    |> expandDecodedRun

decodeLengthAndChar: String -> (Int, String)
decodeLengthAndChar run =
  run
    |> String.toList
    |> List.partition (Char.isDigit)
    |> Tuple.mapFirst convertRunLength
    |> Tuple.mapSecond String.fromList

expandDecodedRun: (Int, String) -> String
expandDecodedRun decodedRun =
  String.repeat (Tuple.first decodedRun) (Tuple.second decodedRun)

convertRunLength: List Char -> Int
convertRunLength runLength =
  runLength
    |> String.fromList
    |> String.toInt
    |> Result.withDefault 1

calculateRunLength: String -> String
calculateRunLength run =
  let runLength = String.length run in
    if runLength == 1 then "" else toString runLength

previousRun: List String -> String
previousRun runs =
  runs
    |> List.head
    |> Maybe.withDefault ""

restOfRuns: List String -> List String
restOfRuns runs =
  runs
    |> List.tail
    |> Maybe.withDefault []

version : Int
version =
  2
