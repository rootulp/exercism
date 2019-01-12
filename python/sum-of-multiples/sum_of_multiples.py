def sum_of_multiples(limit, factors):
    return sum(all_multiples(limit, factors))


def all_multiples(limit, factors):
    multiples = set()
    for factor in factors:
        multiples = multiples.union(get_multiples(limit, factor))
    return multiples

def get_multiples(limit, factor):
    if factor == 0:
        return []

    return [multiple for multiple in range(limit) if multiple % factor == 0]

