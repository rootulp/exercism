class CheckBrackets:
    BRACKETS = {'{': '}',
               '[': ']',
               '(': ')'}
    OPENERS = set(BRACKETS.keys())
    CLOSERS = set(BRACKETS.values())

    def is_paired(self, inp):
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
        return cls.BRACKETS[opener] == closer


def is_paired(inp):
    return CheckBrackets().is_paired(inp)
