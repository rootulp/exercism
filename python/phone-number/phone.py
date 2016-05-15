class Phone:

    INVALID_NUM = "0" * 10
    AREA_CODE_END = 3
    EXCHANGE_CODE_END = 6

    def __init__(self, inp):
        self.number = self.number(inp)

    def number(self, inp):
        cleaned = self.strip(inp)
        if self.valid_11_digits(cleaned):
            return cleaned[1:]
        elif self.valid_10_digits(cleaned):
            return cleaned
        return self.INVALID_NUM

    def area_code(self):
        return self.number[:self.AREA_CODE_END]

    def exchange_code(self):
        return self.number[self.AREA_CODE_END:self.EXCHANGE_CODE_END]

    def subscriber_code(self):
        return self.number[self.EXCHANGE_CODE_END:]

    def pretty(self):
        return "({}) {}-{}".format(self.area_code(), self.exchange_code(),
                                   self.subscriber_code())

    @staticmethod
    def valid_11_digits(inp):
        return len(inp) == 11 and inp.startswith("1")

    @staticmethod
    def valid_10_digits(inp):
        return len(inp) == 10

    @staticmethod
    def strip(inp):
        return ''.join(char for char in inp if char.isdigit())
