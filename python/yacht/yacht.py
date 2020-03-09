"""
This exercise stub and the test suite contain several enumerated constants.

Since Python 2 does not have the enum module, the idiomatic way to write
enumerated constants has traditionally been a NAME assigned to an arbitrary,
but unique value. An integer is traditionally used because it’s memory
efficient.
It is a common practice to export both constants and functions that work with
those constants (ex. the constants in the os, subprocess and re modules).

You can learn more here: https://en.wikipedia.org/wiki/Enumerated_type
"""


# Score categories.
# Change the values as you see fit.
YACHT = 0
FULL_HOUSE = 1
FOUR_OF_A_KIND = 2
LITTLE_STRAIGHT = 3
BIG_STRAIGHT = 4
CHOICE = 5
ONES = 11
TWOS = 12
THREES = 13
FOURS = 14
FIVES = 15
SIXES = 16

def full_house(dice):
    trio = False
    duo = False
    values = set(dice)
    
    for v in values:
        if dice.count(v) == 3:
            trio = True
        if dice.count(v) == 2:
            duo = True

    result = 0
    if trio and duo:
        result = sum(dice)

    return result

def four_of_a_kind(dice):
    values = set(dice)
    result = 0

    for v in values:
        if dice.count(v) >= 4:
            result += v * 4

    return result

def little_straight(dice):
    result = 0
    dice.sort()
    if dice == [1, 2, 3, 4, 5]:
        result = 30

    return result

def big_straight(dice):
    result = 0
    dice.sort()
    if dice == [2, 3, 4, 5, 6]:
        result = 30

    return result

def yacht(dice):
    result = 0
    if len(set(dice)) == 1:
        result = 50
    
    return result


def score(dice, category):
    result = 0

    if category > 10:
        result += (category - 10) * dice.count(category - 10)
    elif category == 0:
        result += yacht(dice)
    elif category == 1:
        result += full_house(dice)
    elif category == 2:
        result += four_of_a_kind(dice)
    elif category == 3:
        result += little_straight(dice)
    elif category == 4:
        result += big_straight(dice)
    elif category == 5:
        result += sum(dice)

    return result
