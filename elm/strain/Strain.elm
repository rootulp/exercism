module Strain exposing (..)

type alias Predicate comparable = (comparable -> Bool)

keep: Predicate comparable -> List comparable -> List comparable
keep predicate collection =
    case collection of
        [] -> []
        comparable::rest ->
            if predicate comparable then
                comparable :: keep predicate rest
            else
                keep predicate rest

discard: Predicate comparable -> List comparable -> List comparable
discard predicate collection =
    List.filter predicate collection
