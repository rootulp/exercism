class Hexa:
    CHAR_VALUES = {'a': 10,
                   'b': 11,
                   'c': 12,
                   'd': 13,
                   'e': 14,
                   'f': 15}
    VALID_CHARS = set(map(str, range(0, 10)) + CHAR_VALUES.keys())

    @classmethod
    def convert(cls, inp):
        if not cls.valid(inp):
            raise ValueError
        return cls.convert_helper(inp)

    @classmethod
    def convert_helper(cls, inp):
        total = 0
        for index, char in enumerate(reversed(inp)):
            total += cls.convert_char(char) * 16**index
        return total

    @classmethod
    def valid(cls, inp):
        return set(inp) <= cls.VALID_CHARS

    @classmethod
    def convert_char(cls, char):
        return int(char) if char.isdigit() else cls.CHAR_VALUES[char]


def hexa(inp):
    return Hexa.convert(inp.lower())
