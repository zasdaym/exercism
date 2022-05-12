import std/[algorithm, sequtils]

type
  Student = tuple
    name: string
    grade: int

  School* = object
    students*: seq[Student]

proc sortedRoster(students: seq[Student]): seq[Student] =
  students.sortedByIt((it.grade, it.name))

proc roster*(s: School): seq[string] =
  sortedRoster(s.students).mapIt(it.name)

proc grade*(s: School, grade: int): seq[string] =
  sorted(s.students.filterIt(it.grade == grade).mapIt(it.name))
