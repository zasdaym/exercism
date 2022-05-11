import sequtils

proc distance*(a, b: string): int =
  if len(a) != len(b):
    raise newException(ValueError, "both strands must have same length")

  result = 0
  for chars in zip(a, b):
    if chars[0] != chars[1]:
      result += 1
