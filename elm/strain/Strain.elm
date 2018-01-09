module Strain exposing (..)

type alias Predicate comparable = (comparable -> Bool)

keep: Predicate comparable -> List comparable -> List comparable
keep predicate collection =
    List.filter predicate collection


discard: Predicate comparable -> List comparable -> List comparable
discard predicate collection =
    List.filter predicate collection
