class Say:

    VALS = {
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
        self.words = 'zero' if num == 0 else self.wordify(num)

    def wordify(self, num):
        self.check_valid(num)
        return ' '.join([self.wordify_chunk(
            i_chunk[1], i_chunk[0]) for i_chunk in enumerate(self.chunkify(num))]).rstrip()

    def chunkify(self, num):
        rev = str(num)[::-1]
        rev_chunks = ([rev[i:i + 3] for i in range(0, len(rev), 3)])[::-1]
        return [int(x[::-1]) for x in rev_chunks]

    def wordify_chunk(self, chunk, i):
        hundreds_digit, left_over = divmod(chunk, 100)
        hundreds = self.get_val(hundreds_digit)

        tens_digit, ones_digit = divmod(left_over, 10)
        if left_over > 10 and left_over < 20:
            tens = self.get_val(left_over)
            ones = None
        else:
            tens = self.get_val(tens_digit * 10)
            ones = self.get_val(ones_digit)

        word_chunk = self.frmt_chunk(hundreds, tens, ones)
        units = self.get_units(len(self.chunkify(self.num)) - 1 - i)
        return word_chunk + ' ' + units if word_chunk else ''

    def frmt_chunk(self, hundreds, tens, ones):
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
        return self.VALS[d]


def say(num):
    return Say(int(num)).words
