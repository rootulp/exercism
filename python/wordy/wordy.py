class Calculator:

    OPERATORS = {"plus": "+",
                 "minus": "-",
                 "multiplied by": "*",
                 "divided by": "/",
                 "What is ": "",
                 "?": ""}

    VALID_TOKENS = set(OPERATORS.values())

    def __init__(self, inp):
        self.inp = inp
        self.tokenized = self.tokenize(inp)
        self.tokens = self.tokenized.split(" ")

    def calculate(self):
        if not self.valid():
            raise ValueError
        return eval(self.tokenized)

    def valid(self):
        return (self.valid_elements() and
                not self.consecutive_tokens() and
                not self.consecutive_digits())

    def consecutive_tokens(self):
        return any(i in self.OPERATORS.values() and
                   j in self.OPERATORS.values() for i, j
                   in self.slices_of_two())

    def consecutive_digits(self):
        return any(self.digit(i) and self.digit(j) for i, j in
                   self.slices_of_two())

    def slices_of_two(self):
        return zip(self.tokens, self.tokens[1:])

    def valid_elements(self):
        return all(self.valid_element(element) for element in self.tokens)

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
    return Calculator(inp).calculate()
