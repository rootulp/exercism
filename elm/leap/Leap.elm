module Leap exposing (..)


isLeapYear : Int -> Bool
isLeapYear year =
    if divisibleBy400 year then
        True
    else if divisibleBy100 year then
        False
    else if divisibleBy4 year then
        True
    else
        False


divisibleBy400 : Int -> Bool
divisibleBy400 year =
    year % 400 == 0


divisibleBy100 : Int -> Bool
divisibleBy100 year =
    year % 100 == 0


divisibleBy4 : Int -> Bool
divisibleBy4 year =
    year % 4 == 0
