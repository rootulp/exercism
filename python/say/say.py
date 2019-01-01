class Say:

    NUM_TO_WORD = {
        1000000000: 'billion',
        1000000: 'million',
        1000: 'thousand',
        90: 'ninety',
        80: 'eighty',
        70: 'seventy',
        60: 'sixty',
        50: 'fifty',
        40: 'forty',
        30: 'thirty',
        20: 'twenty',
        19: 'nineteen',
        17: 'seventeen',
        16: 'sixteen',
        15: 'fifteen',
        14: 'fourteen',
        13: 'thirteen',
        12: 'twelve',
        11: 'eleven',
        10: 'ten',
        9: 'nine',
        8: 'eight',
        7: 'seven',
        6: 'six',
        5: 'five',
        4: 'four',
        3: 'three',
        2: 'two',
        1: 'one',
        0: ''
    }

    def __init__(self, num):
        self.num = num
        self._words = self.get_words(num)

    def get_words(self, num):
        self.check_valid(num)

        if num == 0:
            return 'zero'
        else:
            return ' '.join([self.convert_chunk_to_word(
            chunk, index) for index, chunk in enumerate(self.get_chunks(num))]).rstrip()

    def get_chunks(self, num):
        reversed_string = str(num)[::-1]
        reversed_chunks = ([reversed_string[i:i + 3] for i in range(0, len(reversed_string), 3)])[::-1]
        return [int(x[::-1]) for x in reversed_chunks]

    def convert_chunk_to_word(self, chunk, i):
        hundreds_digit, left_over = divmod(chunk, 100)
        hundreds = self.get_val(hundreds_digit)

        tens_digit, ones_digit = divmod(left_over, 10)
        if 10 < left_over < 20:
            tens = self.get_val(left_over)
            ones = None
        else:
            tens = self.get_val(tens_digit * 10)
            ones = self.get_val(ones_digit)

        word_chunk = self.format_chunk(hundreds, tens, ones)
        units = self.get_units(len(self.get_chunks(self.num)) - 1 - i)
        return word_chunk + ' ' + units if word_chunk else ''

    def format_chunk(self, hundreds, tens, ones):
        chunk = ''
        if hundreds:
            chunk += hundreds + ' hundred '
        if hundreds and tens:
            chunk += 'and '
        if tens:
            chunk += tens
        if tens and ones:
            chunk += '-'
        if ones:
            chunk += ones
        return chunk

    def check_valid(self, num):
        if num < 0 or num > 999999999999:
            raise AttributeError

    def get_units(self, d):
        return self.get_val(1000 ** d) if 1000 ** d > 1 else ''

    def get_val(self, d):
        return self.NUM_TO_WORD[d]

    def in_english(self):
        return self._words


def say(num):
    return Say(int(num)).in_english()
