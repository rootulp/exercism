from itertools import groupby
import re


def encode(s):
    return ''.join([helper(g) for g in [list(g) for k, g in groupby(s)]])


def helper(g):
    return g[0] if len(g) == 1 else str(len(g)) + g[0]


def decode(s):
    groups = re.findall(r'(\d*\D{1})', s)
    pairs = [[re.match(r'\d*', g).group(), g[-1]] for g in groups]

    # Fix hardcoded 0 and 1 indices
    # Also change name of x variable
    return ''.join([int(x[0]) * x[1] if x[0].isdigit() else x[1] for x in pairs])
