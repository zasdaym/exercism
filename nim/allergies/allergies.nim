import bitops, sequtils, sugar, tables

const allergens = {
  "eggs": 1,
  "peanuts": 2,
  "shellfish": 4,
  "strawberries": 8,
  "tomatoes": 16,
  "chocolate": 32,
  "pollen": 64,
  "cats": 128,
}.toOrderedTable

type
  Allergies* = object
    score*: int

proc isAllergicTo*(a: Allergies, allergen: string): bool =
  bitand(a.score, allergens[allergen]) > 0

proc lst*(a: Allergies): seq[string] =
  collect(for n in allergens.keys: n).filterIt(a.isAllergicTo(it))
