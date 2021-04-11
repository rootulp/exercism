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
YACHT = 50
ONES = 1
TWOS = 2
THREES = 3
FOURS = 4
FIVES = 5
SIXES = 6
FULL_HOUSE = None
FOUR_OF_A_KIND = None
LITTLE_STRAIGHT = None
BIG_STRAIGHT = None
CHOICE = None

def score_yacht(dice):
    if len(set(dice)) == 1:
        return YACHT
    else:
        return 0

def score_ones(dice):
    return ONES * dice.count(1)

def score_twos(dice):
    return TWOS * dice.count(2)

def score_threes(dice):
    return THREES * dice.count(3)

def score_fours(dice):
    return FOURS * dice.count(4)

def score_fives(dice):
    return FIVES * dice.count(5)

def score_sixes(dice):
    return SIXES * dice.count(6)

def score(dice, category):
    if (category == YACHT):
        return score_yacht(dice)
    if (category == ONES):
        return score_ones(dice)
    if (category == TWOS):
        return score_twos(dice)
    if (category == THREES):
        return score_threes(dice)
    if (category == FOURS):
        return score_fours(dice)
    if (category == FIVES):
        return score_fives(dice)
    if (category == SIXES):
        return score_sixes(dice)
