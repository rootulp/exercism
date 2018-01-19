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

decodeRun: String -> String
decodeRun run =
  let decodedLengthAndChar =
    run
      |> String.toList
      |> List.partition (Char.isDigit)
      |> Tuple.mapFirst convertRunLength
      |> Tuple.mapSecond String.fromList
  in
    expandDecodedRun decodedLengthAndChar

expandDecodedRun: (Int, String) -> String
expandDecodedRun decodedRun =
  String.repeat (Tuple.first decodedRun) (Tuple.second decodedRun)


convertRunLength: List Char -> Int
convertRunLength runLength =
  runLength
    |> String.fromList
    |> String.toInt
    |> Result.withDefault 1


splitEncodedRuns: (Char -> List String -> List String)
splitEncodedRuns char runs =
  let nextChar = String.fromChar char in
    if Char.isUpper char then
      nextChar :: runs
    else
      String.cons char (previousRun runs) :: (restOfRuns runs)


splitRuns: (Char -> List String -> List String)
splitRuns char runs =
  let nextChar = String.fromChar char in
    if String.startsWith nextChar (previousRun runs) then
      String.cons char (previousRun runs) :: (restOfRuns runs)
    else
      nextChar :: runs

encodeRun: String -> String
encodeRun run =
    run
      |> String.left 1
      |> String.append (runLength run)

runLength: String -> String
runLength run =
  if String.length run == 1 then
    ""
  else
    run
      |> String.length
      |> toString

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
