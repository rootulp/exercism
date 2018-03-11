class IsbnVerifier(object):

    def __init__(self, string):
        self.string = string

    def is_valid(self):
        sum_so_far = 0
        for i, c in enumerate(IsbnVerifier.remove_slashes(self.string)):
            sum_so_far += IsbnVerifier.convert_char_to_int(c) * (10 - i)
        return sum_so_far % 11 == 0

    @staticmethod
    def remove_slashes(string):
        return "".join(filter(lambda char: char != "-", string))

    @staticmethod
    def convert_char_to_int(char):
        return int(IsbnVerifier.convert_x_to_ten(char))

    @staticmethod
    def convert_x_to_ten(char):
        return 10 if char == 'X' else char




def verify(isbn):
    return IsbnVerifier(isbn).is_valid()
