import std/sequtils

proc sum*(n: int, divisors: seq[int]): int =
  result = 0
  for i in countup(1, n - 1):
    if any(divisors, proc (x: int): bool = x != 0 and i mod x == 0):
      result += i
