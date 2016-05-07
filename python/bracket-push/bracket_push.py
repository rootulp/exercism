class CheckBrackets:
    OPENERS = {'{': '}',
               '[': ']',
               '(': ')'}
    CLOSERS = set(OPENERS.values())

    def __init__(self, inp):
        self.check_brackets = self.build_stack(inp)

    def build_stack(self, inp):
        stack = []
        for char in list(inp):
            if char in self.OPENERS:
                stack.append(char)
            elif (char in self.CLOSERS and stack and
                  self.corresponding_brackets(stack[-1], char)):
                stack.pop()
            else:
                return False
        return not bool(stack)

    @classmethod
    def corresponding_brackets(cls, opener, closer):
        return cls.OPENERS[opener] == closer


def check_brackets(inp):
    return CheckBrackets(inp).check_brackets
