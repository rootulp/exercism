module RunLengthEncoding exposing (..)
import Maybe

encode: String -> String
encode data =
  data
    |> String.toList
    |> List.foldl splitIntoChunks []
    |> List.reverse
    |> List.map (\chunk -> String.append (toString (String.length chunk)) (String.left 1 chunk))
    |> String.join ""
    -- |> String.fromList

splitIntoChunks: (Char -> List String -> List String)
splitIntoChunks char chunks =
  if String.startsWith (String.fromChar char) (Maybe.withDefault "" (List.head chunks)) then
    String.cons char (Maybe.withDefault "" (List.head chunks)) :: Maybe.withDefault [] (List.tail chunks)
  else
    String.fromChar char :: chunks


decode: String -> String
decode data =
  data


version : Int
version =
  2
