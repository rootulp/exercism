from fractions import gcd
from itertools import combinations


def is_triplet(sides):
    return Triplet(sides).valid()


def primitive_triplets(n):
    return Triplet.primitive_triplets(n)


def triplets_in_range(low, high):
    return Triplet.triplets_in_range(low, high)


class Triplet:
    def __init__(self, sides):
        self.sides = sides
        self.sorted_sides = sorted(sides)
        self.a = self.sorted_sides[0]
        self.b = self.sorted_sides[1]
        self.c = self.sorted_sides[2]

    def valid(self):
        return self.a**2 + self.b**2 == self.c**2

    def valid_primitive(self):
        return self.valid() and (gcd(self.a, self.b) == 1 and
                                 gcd(self.b, self.c) == 1 and
                                 gcd(self.a, self.c) == 1)

    def valid_in_range(self, low, high):
        valid_range = range(low, high + 1)
        return self.valid() and (self.a in valid_range and
                                 self.b in valid_range and
                                 self.c in valid_range)

    @classmethod
    def triplets(cls, n):
        return set([cls.generate_sides(pair) for pair in
                    cls.valid_pairs(n / 2)])

    @classmethod
    def primitive_triplets(cls, n):
        if cls.odd(n):
            raise ValueError
        else:
            return set([triplet for triplet in cls.triplets(n) if
                       cls(triplet).valid_primitive()])

    @classmethod
    def triplets_in_range(cls, low, high):
        return set([triplet for triplet in cls.possible_triplets(low, high + 1)
                    if cls(triplet).valid_in_range(low, high)])

    @classmethod
    def valid_pairs(cls, n):
        return [pair for pair in cls.factor_pairs(n) if cls.valid_pair(pair)]

    @classmethod
    def valid_pair(cls, pair):
        n, m = pair
        return cls.odd(m - n)

    @staticmethod
    def generate_sides(pair):
        n, m = pair
        return tuple(sorted((m**2 - n**2, 2 * m * n, m**2 + n**2)))

    @staticmethod
    def factor_pairs(n):
        return [(i, n / i) for i in range(1, int(n**0.5) + 1) if n % i == 0]

    @staticmethod
    def possible_triplets(low, high):
        return combinations(range(low, high), 3)

    @staticmethod
    def odd(n):
        return n % 2 == 1
