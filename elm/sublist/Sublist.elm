module Sublist exposing (..)


type ListComparison = Equal | Unequal | Sublist | Superlist


sublist: List Int -> List Int -> ListComparison
sublist list1 list2 =
    Equal


version : Int
version =
    2
