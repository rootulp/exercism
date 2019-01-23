class CheckBrackets:
    BRACKETS = {'{': '}',
               '[': ']',
               '(': ')'}
    OPENERS = set(BRACKETS.keys())
    CLOSERS = set(BRACKETS.values())

    @classmethod
    def is_paired(cls, inp):
        stack = []
        for char in list(inp):
            if char in cls.OPENERS:
                stack.append(char)
            elif (char in cls.CLOSERS and stack and
                  cls.corresponding_brackets(stack[-1], char)):
                stack.pop()
            else:
                return False
        return not bool(stack)

    @classmethod
    def corresponding_brackets(cls, opener, closer):
        return cls.BRACKETS[opener] == closer


def is_paired(inp):
    return CheckBrackets.is_paired(inp)
