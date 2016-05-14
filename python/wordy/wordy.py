class Calculator:

    OPERATORS = {"plus": "+",
                 "minus": "-",
                 "multiplied by": "*",
                 "divided by": "/",
                 "What is ": "",
                 "?": ""}

    VALID_TOKENS = set(OPERATORS.values())

    @classmethod
    def calculate(cls, inp):
        inp = cls.tokenize(inp)
        if not cls.valid(inp):
            raise ValueError
        return eval(inp)

    @classmethod
    def valid(cls, inp):
        return (cls.valid_elements(inp) and
                not cls.consecutive_tokens(inp) and
                not cls.consecutive_digits(inp))

    @classmethod
    def consecutive_tokens(cls, inp):
        return any(i in cls.OPERATORS.values() and j in cls.OPERATORS.values()
                   for i, j in cls.slices_of_two(inp))

    @classmethod
    def consecutive_digits(cls, inp):
        return any(cls.digit(i) and cls.digit(j) for i, j in
                   cls.slices_of_two(inp))

    @classmethod
    def slices_of_two(cls, inp):
        return zip(inp.split(" "), inp.split(" ")[1:])

    @classmethod
    def valid_elements(cls, inp):
        return all(cls.valid_element(element) for element in inp.split(" "))

    @classmethod
    def valid_element(cls, element):
        return element in cls.VALID_TOKENS or cls.digit(element)

    @classmethod
    def tokenize(cls, inp):
        for operator, token in cls.OPERATORS.items():
            inp = inp.replace(operator, token)
        return inp

    @staticmethod
    def digit(element):
        return element.lstrip("-").isdigit()


def calculate(inp):
    return Calculator.calculate(inp)
