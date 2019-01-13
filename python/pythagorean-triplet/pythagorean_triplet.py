from math import gcd
from itertools import product



def triplets(n):
    valid_triplets = set()

    potential_triplets = set([generate_sides(pair) for pair in
                valid_pairs(n // 2)])

    for potential_triplet in potential_triplets:
        if valid_sides(*potential_triplet):
            valid_triplets.add(potential_triplet)

    return valid_triplets

def valid_pairs(n):
    return [pair for pair in potential_pairs(n) if valid_pair(pair)]

def valid_pair(pair):
    n, m = pair
    return odd(m - n)

def generate_sides(pair):
    n, m = pair
    a = m ** 2 - n ** 2
    b = 2 * m * n
    c = m ** 2 + n ** 2
    if (a <= 0 or b <= 0 or c <= 0):
        return None
    return tuple(sorted((a, b, c)))

def valid_sides(a, b, c):
    return a ** 2 + b ** 2 == c ** 2

def potential_pairs(n):
    return product(range(n), range(n))

def odd(n):
    return n % 2 == 1

triplets_with_sum_12 = triplets(12)
print("triplets_with_sum_12 {}".format(triplets_with_sum_12))
