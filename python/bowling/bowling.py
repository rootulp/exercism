

class BowlingGame(object):
    def __init__(self):
        self.rolls = []

    def roll(self, pins):
        self.rolls.append(pins)

    def score(self):
        return sum(self.rolls)
