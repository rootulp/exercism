class Hexa:
    CHAR_VALUES = {'a': 10,
                   'b': 11,
                   'c': 12,
                   'd': 13,
                   'e': 14,
                   'f': 15}
    VALID_CHARS = set(list(map(str, list(range(0, 10)))) + list(CHAR_VALUES.keys()))
    BASE = 16

    @classmethod
    def convert(cls, inp):
        if not cls.valid(inp):
            raise ValueError
        return sum([cls.convert_char(char) * cls.BASE**index for index, char in
                    enumerate(reversed(inp))])

    @classmethod
    def valid(cls, inp):
        return set(inp) <= cls.VALID_CHARS

    @classmethod
    def convert_char(cls, char):
        return int(char) if char.isdigit() else cls.CHAR_VALUES[char]


def hexa(inp):
    return Hexa.convert(inp.lower())
