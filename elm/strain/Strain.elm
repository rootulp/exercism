module Strain exposing (..)

type alias Predicate comparable = (comparable -> Bool)

keep: Predicate comparable -> List a -> List a
keep predicate collection =
    collection


discard: Predicate comparable -> List a -> List a
discard predicate collection =
    collection
