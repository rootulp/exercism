import re


class Handshake:

    events = ['wink', 'double blink', 'close your eyes', 'jump']

    def commands(self, inp):
        if not self.valid_inp(inp):
            return []
        return self.commands_for(self.to_num(inp))

    def commands_for(self, num):
        results = []
        for i in (range(0, 4)):
            if self.testBit(num, i):
                results.append(self.events[i])
        if self.testBit(num, 4):
            results.reverse()
        return results

    def code(self, handshake):
        curr = 0
        for i in handshake:
            if i not in self.events:
                return '0'
            else:
                curr = self.setBit(curr, self.events.index(i))
        if self.events.index(handshake[0]) > self.events.index(handshake[-1]):
            curr = self.setBit(curr, 4)
        return '{0:b}'.format(curr)

    def to_num(self, inp):
        if isinstance(inp, int):
            return inp
        elif isinstance(inp, str):
            return int(inp, 2)

    def valid_inp(self, inp):
        if isinstance(inp, int):
            return self.valid_int(inp)
        elif isinstance(inp, str):
            return self.valid_str(inp)

    def valid_int(self, i):
        return i >= 0

    def valid_str(self, s):
        return not bool(re.search('[^01]', s))

    def valid_handshake(self, handshake):
        return set(self.events) <= set(handshake)

    def testBit(self, int_type, offset):
        mask = 1 << offset
        return (int_type & mask) > 0

    def setBit(self, int_type, offset):
        mask = 1 << offset
        return (int_type | mask)


def handshake(num):
    return Handshake().commands(num)


def code(arr):
    return Handshake().code(arr)
