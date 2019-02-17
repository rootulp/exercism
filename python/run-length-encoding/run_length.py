from itertools import groupby
import re


def encode(string):
    return ''.join([helper(g) for g in [list(group) for _, group in groupby(string)]])


def helper(g):
    return g[0] if len(g) == 1 else str(len(g)) + g[0]


def decode(string):
    groups = re.findall(r'(\d*\D{1})', string)
    pairs = [[re.match(r'\d*', g).group(), g[-1]] for g in groups]

    # Fix hardcoded 0 and 1 indices
    # Also change name of x variable
    return ''.join([int(x[0]) * x[1] if x[0].isdigit() else x[1]
                    for x in pairs])
