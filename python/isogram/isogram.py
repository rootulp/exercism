def is_isogram(word):
    stripped = remove_punctuation(word)
    return len(stripped) == len(set(stripped))

def remove_punctuation(word):
    return filter(str.isalpha, word)
