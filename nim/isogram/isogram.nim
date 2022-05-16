import std/[sequtils, strutils]

proc isIsogram*(s: string): bool =
  let letters = s.toLower.filter(isAlphaAscii)
  letters == letters.deduplicate
