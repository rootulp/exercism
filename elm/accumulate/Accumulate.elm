module Accumulate exposing (..)

accumulate: (a -> b) -> List a -> List b
accumulate mapper collection =
    case collection of

        first :: rest ->
            mapper first :: accumulate mapper rest

        [] ->
            []

