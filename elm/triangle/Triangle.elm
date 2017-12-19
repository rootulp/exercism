module Triangle exposing (..)


type Triangle = Equilateral | Isosceles | Scalene


triangleKind: a -> b -> c -> Result String Triangle
triangleKind a b c =
    (Ok Equilateral)


version : Int
version =
    2
