from itertools import chain


def sum_of_multiples(limit, factors):
    return sum(all_multiples(limit, factors))


def all_multiples(limit, factors):
    multiples = [get_multiples(limit, factor) for factor in factors]
    return set(list(chain(*multiples)))  # remove duplicates


def get_multiples(limit, factor):
    if factor == 0:
        return []
    return [multiple for multiple in range(limit) if multiple % factor == 0]
