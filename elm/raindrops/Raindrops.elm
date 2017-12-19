module Raindrops exposing (..)

raindrops: Int -> String
raindrops number =
    if isFactor 3 number then
        "Pling"
    else if isFactor 5 number then
        "Plang"
    else if isFactor 7 number then
        "Plong"
    else
        toString number

isFactor: Int -> Int -> Bool
isFactor potentialFactor number =
    number % potentialFactor == 0
