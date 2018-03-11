class IsbnVerifier(object):

    VALID_CHARACTERS = set(["-", "X", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"])

    def __init__(self, string):
        self.string = string

    def is_valid(self):
        if self.invalid():
            return False
        sum_so_far = 0
        for i, c in enumerate(self.remove_slashes(self.string)):
            sum_so_far += IsbnVerifier.convert_char_to_int(c) * (10 - i)
        return sum_so_far % 11 == 0

    def invalid(self):
        return self.invalid_character() or self.invalid_length()

    def invalid_character(self):
        return any(char not in self.VALID_CHARACTERS for char in self.string)

    def invalid_length(self):
        return len(self.remove_invalid_characters(self.string)) != 10

    def remove_invalid_characters(self, string):
        return "".join(filter(lambda char: char in self.VALID_CHARACTERS, string))

    def remove_slashes(self, string):
        return "".join(filter(lambda char: char != "-" , string))

    @staticmethod
    def convert_char_to_int(char):
        return int(IsbnVerifier.convert_x_to_ten(char))

    @staticmethod
    def convert_x_to_ten(char):
        return 10 if char == 'X' else char




def verify(isbn):
    return IsbnVerifier(isbn).is_valid()
