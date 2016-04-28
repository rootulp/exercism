class Binary:
    VALID_CHARS = set("01")

    @classmethod
    def parse_binary(cls, inp):
        if not cls.valid(inp):
            raise ValueError
        return cls.convert_to_decimal(inp)

    @classmethod
    def convert_to_decimal(cls, inp):
        return sum([2**idx for idx, val in enumerate(reversed(inp))
                    if val == "1"])

    @classmethod
    def valid(cls, inp):
        return set(inp) <= cls.VALID_CHARS


def parse_binary(inp):
    return Binary.parse_binary(inp)
