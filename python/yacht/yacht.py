from collections import Counter

# Score categories
ONES = "ONES"
TWOS = "TWOS"
THREES = "THREES"
FOURS = "FOURS"
FIVES = "FIVES"
SIXES = "SIXES"
FULL_HOUSE = "FULL_HOUSE"
FOUR_OF_A_KIND = "FOUR_OF_A_KIND"
STRAIGHT = "STRAIGHT"
LITTLE_STRAIGHT = "LITTLE_STRAIGHT"
BIG_STRAIGHT = "BIG_STRAIGHT"
CHOICE = "CHOICE"
YACHT = "YACHT"

# Score per category
scores = {
    ONES: 1,
    TWOS: 2,
    THREES: 3,
    FOURS: 4,
    FIVES: 5,
    SIXES: 6,
    STRAIGHT: 30,
    YACHT: 50,
}

def score_ones(dice):
    return scores[ONES] * dice.count(1)

def score_twos(dice):
    return scores[TWOS] * dice.count(2)

def score_threes(dice):
    return scores[THREES] * dice.count(3)

def score_fours(dice):
    return scores[FOURS] * dice.count(4)

def score_fives(dice):
    return scores[FIVES] * dice.count(5)

def score_sixes(dice):
    return scores[SIXES] * dice.count(6)

def score_full_house(dice):
    return sum(dice) if sorted(set(Counter(dice).values())) == [2, 3] else 0

def score_four_of_a_kind(dice):
    counter = Counter(dice)
    if 4 in set(counter.values()):
        for k, v in counter.items():
            if v == 4:
                return k * 4
    if 5 in set(counter.values()):
        return dice[0] * 4
    return 0

def score_little_straight(dice):
    return scores[STRAIGHT] if sorted(dice) == [1, 2, 3, 4, 5] else 0

def score_big_straight(dice):
    return scores[STRAIGHT] if sorted(dice) == [2, 3, 4, 5, 6] else 0

def score_choice(dice):
    return sum(dice)

def score_yacht(dice):
    return scores[YACHT] if len(set(dice)) == 1 else 0

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
    if (category == FOUR_OF_A_KIND):
        return score_four_of_a_kind(dice)
    if (category == LITTLE_STRAIGHT):
        return score_little_straight(dice)
    if (category == BIG_STRAIGHT):
        return score_big_straight(dice)
    if (category == CHOICE):
        return score_choice(dice)
    if (category == YACHT):
        return score_yacht(dice)
