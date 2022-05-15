import std/[sequtils, sugar]

proc sum*(n: int, divisors: seq[int]): int =
  result = 0
  for i in 1..n-1:
    if any(divisors, x => x != 0 and i mod x == 0):
      result += i
