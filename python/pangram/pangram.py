ALPHABET = 'abcdefghijklmnopqrstuvwxyz '


def is_pangram(s):
    return set(list(s.lower())) >= set(ALPHABET)
