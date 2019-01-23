def detect_anagrams(w, poss):
    return filter(lambda p: anagram(w, p), poss);

def anagram(w, p):
    return w.lower() != p.lower() and sort(w.lower()) == sort(p.lower())

def sort(word):
    return sorted(list(word.lower()));