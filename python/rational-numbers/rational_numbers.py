from __future__ import division
from math import gcd

class Rational:
    def __init__(self, numer, denom):
        if denom < 0:
            numer, denom, = -numer, -denom

        cdiv = gcd(numer, denom)
        self.numer = numer // cdiv
        self.denom = denom // cdiv

    def __eq__(self, other):
        return self.numer == other.numer and self.denom == other.denom

    def __repr__(self):
        return '{}/{}'.format(self.numer, self.denom)

    def __add__(self, other):
        a1, b1, a2, b2 = self.__expand(other)
        return Rational(a1 * b2 + a2 * b1, b1 * b2)

    def __sub__(self, other):
        a1, b1, a2, b2 = self.__expand(other)
        return Rational(a1 * b2 - a2 * b1, b1 * b2) 

    def __mul__(self, other):
        a1, b1, a2, b2 = self.__expand(other)
        return Rational(a1 * a2, b1 * b2)

    def __truediv__(self, other):
        a1, b1, a2, b2 = self.__expand(other)
        return Rational(a1 * b2, a2 * b1)

    def __abs__(self):
        return Rational(abs(self.numer), abs(self.denom))

    def __pow__(self, power):
        p = power
        if power < 0:
            p = abs(power)
            result = Rational(self.denom ** p, self.numer ** p)
        else:
            result = Rational(self.numer ** p, self.denom ** p)
        
        return result

    def __rpow__(self, base):
        return base ** (self.numer / self.denom)

    def __expand(self, other):
        return self.numer, self.denom, other.numer, other.denom
