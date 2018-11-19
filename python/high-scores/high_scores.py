import heapq

class HighScores(object):
    def __init__(self, scores):
        self.scores = scores

    def highest(self):
        return max(self.scores)

    def latest(self):
        return self.scores[-1]

    def report(self):
        return f"Your latest score was {self.latest()}. That's {self.amount_short()}"

    def amount_short(self):
        if self.highest() == self.latest():
            return "your personal best!"
        else:
            return f"{self.highest() - self.latest()} short of your personal best!"

    def top(self):
        return heapq.nlargest(3, self.scores)
