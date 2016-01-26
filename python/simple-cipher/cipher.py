import random
import string


class Caesar:
    @classmethod
    def encode(cls, phrase):
        return ''.join([chr(cls.wrap(ord(c) + 3)) for c in cls.clean(phrase)])

    @classmethod
    def decode(cls, phrase):
        return ''.join([chr(cls.wrap(ord(c) - 3)) for c in cls.clean(phrase)])

    @classmethod
    def wrap(cls, val):
        if val > 122:
            val -= 26
        elif val < 97:
            val += 26
        return val

    @classmethod
    def clean(cls, phrase):
        return filter(str.isalpha, phrase.lower())


class Cipher:

    def __init__(self, key=None):
        self.random_key_length = 100
        self.key = key if key is not None else self.generate_random_key()
        if not self.valid_key():
            raise ValueError()

    def encode(self, phrase):
        return ''.join([chr(self.wrap(ord(c) + self.offset(i))) for i, c in
                        enumerate(self.clean(phrase))])

    def decode(self, phrase):
        return ''.join([chr(self.wrap(ord(c) - self.offset(i))) for i, c in
                        enumerate(self.clean(phrase))])

    def wrap(self, val):
        while val > 122:
            val -= 26
        while val < 97:
            val += 26
        return val

    def clean(self, phrase):
        return filter(str.isalpha, phrase.lower())

    def offset(self, index):
        return self.wrap(ord(self.key[index % len(self.key)]) - 97)

    def generate_random_key(self):
        return ''.join(random.SystemRandom().choice(string.ascii_lowercase)
                       for _ in range(self.random_key_length))

    def valid_key(self):
        return self.key.isalpha() and self.key.islower()
