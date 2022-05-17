import std/[math, sequtils]

proc isValid(sides: array[0..2, int]): bool =
  if 0 in sides:
    return false

  let hypotenuse = sides[maxIndex(sides)]
  hypotenuse <= sum(sides) - hypotenuse

proc isEquilateral*(sides: array[0..2, int]): bool =
  isValid(sides) and len(sides.deduplicate) == 1

proc isIsosceles*(sides: array[0..2, int]): bool =
  isValid(sides) and len(sides.deduplicate) < 3

proc isScalene*(sides: array[0..2, int]): bool =
  isValid(sides) and len(sides.deduplicate) == 3
