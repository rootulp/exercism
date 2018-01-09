module Strain exposing (..)


keep: (comparable -> Bool) -> List a -> List a
keep predicate collection =
    collection


discard: (comparable -> Bool) -> List a -> List a
discard predicate collection =
    collection
