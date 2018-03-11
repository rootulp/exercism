class IsbnVerifier(object):

    VALID_CHARACTERS = set(["-", "X", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"])
    VALID_LENGTH = 10

    @classmethod
    def is_valid(cls, string):
        if cls.invalid(string):
            return False
        sum_so_far = 0
        for i, c in enumerate(cls.remove_slashes(string)):
            sum_so_far += cls.convert_char_to_int(c) * (10 - i)
        return sum_so_far % 11 == 0

    @classmethod
    def invalid(cls, string):
        return cls.invalid_character(string) or cls.invalid_length(string)

    @classmethod
    def invalid_character(cls, string):
        return any(char not in cls.VALID_CHARACTERS for char in string)

    @classmethod
    def invalid_length(cls, string):
        return len(cls.remove_invalid_characters(string)) != cls.VALID_LENGTH

    @classmethod
    def remove_invalid_characters(cls, string):
        return "".join(filter(lambda char: char in cls.VALID_CHARACTERS, string))

    @classmethod
    def convert_char_to_int(cls, char):
        return int(cls.convert_x_to_ten(char))

    @classmethod
    def convert_x_to_ten(cls, char):
        return 10 if char == 'X' else char

    @staticmethod
    def remove_slashes(string):
        return "".join(filter(lambda char: char != "-" , string))


def verify(isbn):
    return IsbnVerifier.is_valid(isbn)
