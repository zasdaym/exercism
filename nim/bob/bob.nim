import std/strutils

proc hey*(s: string): string =
  let stripped = strip(s)
  let isEmpty = isEmptyOrWhitespace(stripped)
  let isQuestion = endsWith(stripped, "?")
  let isAllUpper = contains(stripped, Letters) and toUpperAscii(stripped) == stripped

  result =
    if isEmpty: "Fine. Be that way!"
    elif isQuestion and isAllUpper: "Calm down, I know what I'm doing!"
    elif isQuestion: "Sure."
    elif isAllUpper: "Whoa, chill out!"
    else: "Whatever."
