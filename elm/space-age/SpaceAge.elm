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

ageOn: Planet -> Float -> Float
ageOn planet ageInSeconds =
  ageInSeconds / orbitalPeriod planet

orbitalPeriod: Planet -> Float
orbitalPeriod planet =
  case planet of
    Mercury -> 7600530.24
    Venus -> 19413907.2
    Earth -> 31558149.76
    Mars -> 59354294.4
    Jupiter -> 374335776.0
    Saturn -> 929596608.0
    Uranus -> 2661041808.0
    Neptune -> 5200418592.0
