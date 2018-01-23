module Sublist exposing (..)


type ListComparison = Equal | Unequal | Sublist | Superlist


sublist: List Int -> List Int -> ListComparison
sublist list1 list2 =
    if list1 == list2 then
        Equal
    else if List.all (\element -> List.member element list2) list1 then
        Sublist
    else if List.all (\element -> List.member element list1) list2 then
        Superlist
    else
        Unequal


version : Int
version =
    2
