def detect_anagrams(w, poss):
    return [p for p in poss if anagram(w, p)]


def anagram(w, p):
    return w.lower() != p.lower() and sort(w.lower()) == sort(p.lower())


def sort(word):
    return sorted(list(word.lower()))
