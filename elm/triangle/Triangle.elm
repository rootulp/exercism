module Triangle exposing (..)


type Triangle = Equilateral | Isosceles | Scalene


triangleKind: comparable -> comparable -> comparable -> Result String Triangle
triangleKind a b c =
    if invalidLengths a b c then
        (Err "Invalid lengths")
    else if violatesInequality a b c then
        (Err "Violates inequality")
    else if isEquilateral a b c then
        (Ok Equilateral)
    else if isIsosceles a b c then
        (Ok Isosceles)
    else
        (Ok Scalene)


isEquilateral: comparable -> comparable -> comparable -> Bool
isEquilateral a b c =
    a == b  && b == c


isIsosceles: comparable -> comparable -> comparable -> Bool
isIsosceles a b c =
    (a == b && b /= c) ||
    (a == c && c /= b) ||
    (b == c && c /= a)


invalidLengths: comparable -> comparable -> comparable -> Bool
invalidLengths a b c =
    a <= 0 || b <= 0 || c <= 0


violatesInequality: comparable -> comparable -> comparable -> Bool
violatesInequality a b c =
    (a + b <= c) ||
    (a + c <= b) ||
    (b + c <= a)


version : Int
version =
    2
