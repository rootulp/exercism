def find_anagrams(word, candidates):
    return [
        candidate for candidate in candidates if is_anagram(
            word, candidate)]

# Anagram comparison is case insensitive


def is_anagram(word, candidate):
    word_lowercase = word.lower()
    candidate_lowercase = candidate.lower()
    return (is_not_identical(word_lowercase, candidate_lowercase) and
            is_identical_when_sorted(word_lowercase, candidate_lowercase))


def is_not_identical(word, candidate):
    return word != candidate


def is_identical_when_sorted(word, candidate):
    return sorted(word) == sorted(candidate)
