import string
import random


class Robot:

    def __init__(self):
        self.name = self.generate_name()

    def reset(self):
        self.name = self.generate_name()

    def generate_name(self):
        random.seed()
        return self.random_prefix(2) + self.random_suffix(3)

    def random_prefix(self, n):
        return ''.join(random.sample(string.ascii_uppercase, n))

    def random_suffix(self, n):
        return ''.join(random.sample(string.digits, n))
