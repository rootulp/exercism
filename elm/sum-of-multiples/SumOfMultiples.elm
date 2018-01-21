module SumOfMultiples exposing (..)

import Set


sumOfMultiples : List Int -> Int -> Int
sumOfMultiples numbers maxMultiple =
    numbers
        |> List.concatMap (\number -> multiplesOf number maxMultiple [])
        |> removeDuplicates
        |> List.sum


multiplesOf : Int -> Int -> List Int -> List Int
multiplesOf baseNumber maxMultiple multiplesSoFar =
    let
        nextMultiple =
            calculateNextMultiple baseNumber (List.head multiplesSoFar)
    in
    if nextMultiple < maxMultiple then
        multiplesOf baseNumber maxMultiple (nextMultiple :: multiplesSoFar)
    else
        multiplesSoFar


calculateNextMultiple : Int -> Maybe Int -> Int
calculateNextMultiple baseNumber previousMultiple =
    baseNumber + Maybe.withDefault 0 previousMultiple


removeDuplicates : List Int -> List Int
removeDuplicates numbers =
    numbers
        |> Set.fromList
        |> Set.toList
