module Sublist exposing (..)

import Dict exposing (Dict)

type ListComparison = Equal | Unequal | Sublist | Superlist


sublist: List Int -> List Int -> ListComparison
sublist list1 list2 =
    if list1 == list2 then
        Equal
    else if isSublist list1 list2 then
        Sublist
    else if isSublist list2 list1 then
        Superlist
    else
        Unequal


isSublist: List Int -> List Int -> Bool
isSublist list1 list2 =
    let dict1 = occurences list1
        dict2 = occurences list2
    in
        Dict.merge onlyInLeft inBoth onlyInRight dict1 dict2 True


onlyInLeft: (comparable -> Int -> Bool -> Bool)
onlyInLeft data count1 previousResults =
    previousResults && False

inBoth: (comparable -> Int -> Int -> Bool -> Bool)
inBoth data count1 count2 previousResults =
    previousResults && (count1 <= count2)

onlyInRight: (comparable -> Int -> Bool -> Bool)
onlyInRight data count2 previousResults =
    previousResults && True

occurences : List Int -> Dict Int Int
occurences list =
    list
        |> List.foldl populateDict Dict.empty


populateDict : comparable -> Dict comparable Int -> Dict comparable Int
populateDict data dict =
    Dict.update data incrementValue dict


incrementValue : Maybe Int -> Maybe Int
incrementValue maybe =
    Maybe.withDefault 0 maybe
        + 1
        |> Just

version : Int
version =
    2
