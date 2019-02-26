def is_empty(stack):
    return stack == []


class CheckBrackets:
    BRACKETS = {'{': '}',
                '[': ']',
                '(': ')'}
    OPENING_BRACKET = set(BRACKETS.keys())
    CLOSING_BRACKET = set(BRACKETS.values())

    def __init__(self, inp):
        self.inp = inp

    def is_paired(self):
        stack = []
        for bracket in self.get_brackets(self.inp):
            if self.is_opening_bracket(bracket):
                stack.append(bracket)
            elif self.is_closing_bracket(bracket) and self.closes_existing_bracket(bracket, stack):
                stack.pop()
            else:
                return False  # This is an invalid closing bracket

        return is_empty(stack)  # There are more open brackets left to close

    def closes_existing_bracket(self, char, stack):
        return stack and self.matching_brackets(stack[-1], char)

    def matching_brackets(self, opener, closer):
        return self.BRACKETS[opener] == closer

    def get_brackets(self, string):
        return [char for char in string if self.is_bracket(char)]

    def is_opening_bracket(self, bracket):
        return bracket in self.OPENING_BRACKET

    def is_closing_bracket(self, bracket):
        return bracket in self.CLOSING_BRACKET

    def is_bracket(self, char):
        return self.is_opening_bracket(char) or self.is_closing_bracket(char)


def is_paired(inp):
    return CheckBrackets(inp).is_paired()
