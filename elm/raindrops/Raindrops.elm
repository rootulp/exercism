module Raindrops exposing (..)

raindrops: Int -> String
raindrops number =
    if isFactor 3 number || isFactor 5 number || isFactor 7 number then
        plingPlangPlong number
    else
        toString number

plingPlangPlong: Int -> String
plingPlangPlong number =
    String.join "" [
        if isFactor 3 number then "Pling" else "",
        if isFactor 5 number then "Plang" else "",
        if isFactor 7 number then "Plong" else ""
    ]

isFactor: Int -> Int -> Bool
isFactor potentialFactor number =
    number % potentialFactor == 0
