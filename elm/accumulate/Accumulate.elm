module Accumulate exposing (..)

-- Not a real soltuion because this uses List.map
accumulate: (a -> b) -> List a -> List b
accumulate mapper collection =
    List.map mapper collection
