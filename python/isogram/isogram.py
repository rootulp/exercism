def is_isogram(word):
    return len(remove_punctuation(word)) == len(set(remove_punctuation(word)))


def remove_punctuation(word):
    return list(filter(str.isalpha, word.lower()))
