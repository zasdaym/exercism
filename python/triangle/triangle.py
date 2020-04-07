def equilateral(sides):
    ok, uniqueness = checkTriangle(sides)
    return uniqueness == 1


def isosceles(sides):
    ok, uniqueness = checkTriangle(sides)
    return ok and uniqueness < 3


def scalene(sides):
    ok, uniqueness = checkTriangle(sides)
    return ok and uniqueness == 3


def checkTriangle(sides):
    if 0 in sides:
        return False, -1

    sides.sort()
    return sides[0] + sides[1] >= sides[2], len(set(sides))
