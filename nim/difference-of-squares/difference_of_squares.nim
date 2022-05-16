proc squareOfSum*(n: int): int =
  let sum = (1 + n) * n div 2
  sum * sum

proc sumOfSquares*(n: int): int =
  (n * (n + 1) * (2 * n + 1)) div 6

proc difference*(n: int): int =
  squareOfSum(n) - sumOfSquares(n)
