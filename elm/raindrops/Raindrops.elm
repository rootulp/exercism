module Raindrops exposing (..)

raindrops: Int -> String
raindrops number =
    if isFactor 3 number then
        "Pling"
    else
        toString number

isFactor: Int -> Int -> Bool
isFactor potentialFactor number =
    number % potentialFactor == 0
