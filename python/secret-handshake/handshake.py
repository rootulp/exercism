import re


class Handshake:

    EVENTS = ['wink', 'double blink', 'close your eyes', 'jump']

    def commands(self, inp):
        if not self.valid_inp(inp):
            return []
        return self.commands_for_num(self.to_num(inp))

    def commands_for_num(self, num):
        if self.testBit(num, 4):
            return list(reversed(self.unreversed_commands(num)))
        return self.unreversed_commands(num)

    def unreversed_commands(self, num):
        return [self.EVENTS[bit] for bit in range(0, 4) if
                self.testBit(num, bit)]

    def code(self, handshake):
        if not self.valid_handshake(handshake):
            return '0'
        return self.code_for_handshake(handshake)

    def code_for_handshake(self, handshake):
        code = self.unreversed_code(handshake)
        if self.EVENTS.index(handshake[0]) > self.EVENTS.index(handshake[-1]):
            # Prepend code with 1 or add 'reverse' bit
            code = self.setBit(code, 4)
        return '{0:b}'.format(code)

    def unreversed_code(self, handshake):
        curr = 0
        for event in handshake:
            curr = self.setBit(curr, self.EVENTS.index(event))
        return curr

    def valid_inp(self, inp):
        if isinstance(inp, int):
            return self.valid_integer(inp)
        elif isinstance(inp, str):
            return self.valid_string(inp)

    def valid_handshake(self, handshake):
        return set(handshake) <= set(self.EVENTS)

    @staticmethod
    def valid_integer(integer):
        return integer >= 0

    @staticmethod
    def valid_string(string):
        return not bool(re.search('[^01]', string))

    @staticmethod
    def testBit(int_type, offset):
        mask = 1 << offset
        return (int_type & mask) > 0

    @staticmethod
    def setBit(int_type, offset):
        mask = 1 << offset
        return (int_type | mask)

    @staticmethod
    def to_num(inp):
        if isinstance(inp, str):
            return int(inp, 2)
        return inp


def handshake(num):
    return Handshake().commands(num)


def code(arr):
    return Handshake().code(arr)
