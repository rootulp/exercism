module SpaceAge exposing (..)

type Planet
  = Mercury
  | Venus
  | Earth
  | Mars
  | Jupiter
  | Saturn
  | Uranus
  | Neptune

ageOn : Planet -> Int -> Float
ageOn planet ageInSeconds =
  32
