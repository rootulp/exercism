module RunLengthEncoding exposing (..)
import Maybe

encode: String -> String
encode data =
  data
    |> String.toList
    |> List.foldr splitRuns []
    |> List.map encodeRun
    |> String.join ""

encodeRun: String -> String
encodeRun run =
  String.append (toString (String.length run)) (String.left 1 run)


splitRuns: (Char -> List String -> List String)
splitRuns char runs =
  if String.startsWith (String.fromChar char) (previousRun runs) then
    String.cons char (previousRun runs) :: (restOfRuns runs)
  else
    String.fromChar char :: runs

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


decode: String -> String
decode data =
  data


version : Int
version =
  2
