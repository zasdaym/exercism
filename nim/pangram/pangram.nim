import sequtils, strutils

proc isPangram*(s: string): bool =
  {'a'..'z'}.allIt(it in s.toLower)
