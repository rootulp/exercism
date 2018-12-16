from functools import reduce


def sum_of_multiples(limit, factors=[3, 5]):
    mult = set()
    for factor in factors:
        if factor <= 0:
            continue
        for curr_mult in range(1, limit / factor + 1):
            if factor * curr_mult < limit:
                mult.add(factor * curr_mult)

    return 0 if len(mult) == 0 else reduce(lambda x, y: x + y, mult)
