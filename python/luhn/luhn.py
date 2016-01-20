class Luhn:

    def __init__(self, number):
        self.number = number

    def checksum(self):
        return sum(self.addends()) % 10

    def addends(self):
        return [self.addend(idx, int(val)) for idx, val in
                enumerate(reversed(str(self.number)))]

    def addend(self, idx, val):
        return self.subtract_nine(idx, self.double_every_other(idx, val))

    def double_every_other(self, idx, val):
        return val * 2 if idx % 2 == 1 else val

    def subtract_nine(self, idx, val):
        return val - 9 if val > 10 else val

    def is_valid(self):
        return self.checksum() == 0

    @classmethod
    def create(cls, num):
        for i in range(0, 10):
            if cls(int(str(num) + str(i))).is_valid():
                return int(str(num) + str(i))
