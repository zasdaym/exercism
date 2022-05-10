import std/strutils

proc hey*(s: string): string =
  var stripped = strip(s)
  let isEmpty = isEmptyOrWhitespace(stripped)
  let isQuestion = endsWith(stripped, "?")
  let isAllUpper = contains(stripped, Letters) and toUpperAscii(stripped) == stripped

  result = "Whatever."
  if isEmpty:
    result = "Fine. Be that way!"
  if isQuestion and isAllUpper:
    result = "Calm down, I know what I'm doing!"
  elif isQuestion:
    result = "Sure."
  elif isAllUpper:
    result = "Whoa, chill out!"
