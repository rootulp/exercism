module SumOfMultiples exposing (..)
import Set

sumOfMultiples: List Int -> Int -> Int
sumOfMultiples numbers maxMultiple =
    numbers
        |> List.concatMap (\number -> multiplesOf number maxMultiple [])
        |> removeDuplicates
        |> List.sum

multiplesOf: Int -> Int -> List Int -> List Int
multiplesOf baseNumber maxMultiple multiplesSoFar =
    let nextMultiple = Maybe.withDefault 0 (List.head multiplesSoFar) + baseNumber in
    if nextMultiple < maxMultiple then
        multiplesOf baseNumber maxMultiple (nextMultiple :: multiplesSoFar)
    else
        multiplesSoFar

removeDuplicates: List Int -> List Int
removeDuplicates numbers =
    numbers
        |> Set.fromList
        |> Set.toList
