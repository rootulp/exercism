module Sublist exposing (..)


type ListComparison
    = Equal
    | Unequal
    | Sublist
    | Superlist


sublist : List Int -> List Int -> ListComparison
sublist list1 list2 =
    if list1 == list2 then
        Equal
    else if isSublist list1 list2 then
        Sublist
    else if isSublist list2 list1 then
        Superlist
    else
        Unequal


isSublist : List Int -> List Int -> Bool
isSublist list1 list2 =
    if list1 == List.take (List.length list1) list2 then
        True
    else if List.length list1 > List.length list2 then
        False
    else
        isSublist list1 (List.drop 1 list2)


version : Int
version =
    2
