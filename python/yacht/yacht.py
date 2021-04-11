from collections import Counter

# Score categories
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

def score_full_house(dice):
    return sum(dice) if sorted(set(Counter(dice).values())) == [2, 3] else 0

def score_yacht(dice):
    return YACHT if len(set(dice)) == 1 else 0

def score(dice, category):
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
    if (category == FULL_HOUSE):
        return score_full_house(dice)
    if (category == YACHT):
        return score_yacht(dice)
