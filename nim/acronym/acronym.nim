import std/[re, strutils]

let abbrevRegex = re"\b[a-zA-Z]"

proc abbreviate*(s: string): string =
  s.multiReplace(("'", ""), ("_", "")).toUpper.findAll(abbrevRegex).join
