import heapq


class HighScores(object):
    def __init__(self, scores):
        self.scores = scores

    def highest(self):
        return max(self.scores)

    def latest(self):
        return self.scores[-1]

    def report(self):
        return f"{self.latest_score_message()} {self.personal_best_message()}"

    def latest_score_message(self):
        return f"Your latest score was {self.latest()}."

    def personal_best_message(self):
        if self.amount_short() == 0:
            return "That's your personal best!"
        else:
            return f"That's {self.amount_short()} short of your personal best!"

    def amount_short(self):
        return self.highest() - self.latest()

    def top(self):
        return heapq.nlargest(3, self.scores)
