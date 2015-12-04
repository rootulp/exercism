class Say:

    UNITS = {
        3: 'billion',
        2: 'million',
        1: 'thousand',
        0: ''
    }

    TENS = {
        9: 'ninety',
        8: 'eighty',
        7: 'seventy',
        6: 'sixty',
        5: 'fifty',
        4: 'forty',
        3: 'thirty',
        2: 'twenty',
        0: ''
    }

    TEENS = {
        19: 'nineteen',
        17: 'seventeen',
        16: 'sixteen',
        15: 'fifteen',
        14: 'fourteen',
        13: 'thirteen',
        12: 'twelve',
        11: 'eleven',
        10: 'ten'
    }

    DIGITS = {
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
        self.words = self.wordify(num)

    def wordify(self, num):
        if num < 0 or num > 999999999999:
            raise AttributeError
        elif num == 0:
            return 'zero'
        else:
            result = ''
            for i, chunk in enumerate(self.chunkify(num)):
                units = self.get_units(len(self.chunkify(num)) - 1 - i)
                val = self.wordify_chunk(chunk)
                if val:
                    result += val + ' ' + units + ' '

            return result.rstrip()

    def chunkify(self, num):
        rev = str(num)[::-1]
        rev_chunks = ([rev[i:i + 3] for i in range(0, len(rev), 3)])[::-1]
        return map(lambda x: int(x[::-1]), rev_chunks)

    def wordify_chunk(self, chunk):
        hundreds_digit, left = divmod(chunk, 100)
        hundreds = self.get_digit(hundreds_digit)

        tens_digit, ones_digit = divmod(left, 10)
        if left > 10 and left < 20:
            tens = self.get_teen(left)
            ones = None
        else:
            tens = self.get_tens(tens_digit)
            ones = self.get_digit(ones_digit)

        return self.frmt_chunk(hundreds, tens, ones)

    def frmt_chunk(self, hundreds, tens, ones):
        result = ''
        if hundreds:
            result += hundreds + ' hundred '
        if hundreds and tens:
            result += 'and '
        if tens:
            result += tens
        if tens and ones:
            result += '-'
        if ones:
            result += ones
        return result

    def get_digit(self, d):
        return self.DIGITS[d]

    def get_teen(self, d):
        return self.TEENS[d]

    def get_tens(self, d):
        return self.TENS[d]

    def get_units(self, d):
        return self.UNITS[d]


def say(num):
    return Say(int(num)).words
