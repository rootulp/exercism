import string


class PigLatinTranslator:

    alpha = set(string.ascii_lowercase)
    vowels = set(['a', 'e', 'i', 'o', 'u'])
    consonants = alpha - vowels

    @classmethod
    def translate_phrase(cls, phrase):
        return ' '.join([cls.translate(word) for word in phrase.split()])

    @classmethod
    def translate(cls, word):
        if (word[0] in cls.vowels or
            word.startswith('yt') or
                word.startswith('xr')):
            return word + 'ay'
        elif (word.startswith('squ') or
                word.startswith('sch') or
                word.startswith('thr')):
            return word[3:] + word[0:3] + 'ay'
        elif ((word[0] in cls.consonants and
                word[1] in cls.consonants) or
                word.startswith('qu')):
            return word[2:] + word[0:2] + 'ay'
        elif (word[0] in cls.consonants):
            return word[1:] + word[0] + 'ay'


def translate(phrase):
    print((PigLatinTranslator.alpha))
    print((PigLatinTranslator.vowels))
    print((PigLatinTranslator.consonants))
    return PigLatinTranslator.translate_phrase(phrase)
