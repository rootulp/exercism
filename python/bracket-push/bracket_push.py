def isEmpty(stack):
    return stack == []

class CheckBrackets:
    BRACKETS = {'{': '}',
               '[': ']',
               '(': ')'}
    OPENERS = set(BRACKETS.keys())
    CLOSERS = set(BRACKETS.values())

    def __init__(self, inp):
        self.inp = inp

    def is_paired(self):
        stack = []
        for char in list(self.inp):
            if char in self.OPENERS:
                stack.append(char)
            elif self.is_closing_bracket(char, stack):
                stack.pop()
            else:
                return False
        return isEmpty(stack)

    def is_closing_bracket(self, char, stack):
        return char in self.CLOSERS and stack and self.matching_brackets(stack[-1], char)

    def matching_brackets(self, opener, closer):
        return self.BRACKETS[opener] == closer


def is_paired(inp):
    return CheckBrackets(inp).is_paired()
