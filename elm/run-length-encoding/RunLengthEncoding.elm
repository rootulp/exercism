module RunLengthEncoding exposing (..)
import Maybe

encode: String -> String
encode data =
  data
    |> String.toList
    |> List.foldl splitRuns []
    |> List.reverse
    |> List.map (\chunk -> String.append (toString (String.length chunk)) (String.left 1 chunk))
    |> String.join ""

splitRuns: (Char -> List String -> List String)
splitRuns char runs =
  if String.startsWith (String.fromChar char) (lastRun runs) then
    String.cons char (lastRun runs) :: Maybe.withDefault [] (List.tail runs)
  else
    String.fromChar char :: runs

lastRun: List String -> String
lastRun runs =
  Maybe.withDefault "" (List.head runs)


decode: String -> String
decode data =
  data


version : Int
version =
  2
