import math
from functools import reduce


def is_perfect(number):
    return sum(factors(number)) == number


def factors(n):
    return set(reduce(list.__add__, pairs_of_factors(n))) - set([n])


def pairs_of_factors(n):
    return [[i, n / i] for i in xrange(1, int(math.sqrt(n)) + 1) if n % i == 0]
