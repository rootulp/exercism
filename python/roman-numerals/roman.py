class Roman:

    NUMERALS = {
        1: "I",
        4: "IV",
        5: "V",
        9: "IX",
        10: "X",
        40: "XL",
        50: "L",
        90: "XC",
        100: "C",
        400: "CD",
        500: "D",
        900: "CM",
        1000: "M"
    }

    @classmethod
    def numeral(cls, arabic):
        return ''.join([cls.NUMERALS[key] for key in cls.get_components(arabic)])

    @classmethod
    def get_components(cls, arabic):
        components = []
        for key in reversed(sorted(cls.NUMERALS.keys())):
            while arabic >= key:
                arabic -= key
                components.append(key)
        return components


def numeral(arabic):
    return Roman.numeral(arabic)
