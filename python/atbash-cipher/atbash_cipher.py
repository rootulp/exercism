import string


class Atbash:

    PLAIN = 'abcdefghijklmnopqrstuvwxyz'
    PRIME = 'zyxwvutsrqponmlkjihgfedcba'
    CIPHER = dict(zip(list(PLAIN), list(PRIME)))
    EXCLUDE = set(string.punctuation + ' ')

    @staticmethod
    def encode(self, msg):
        return self.split_every_five(self, self.encoded(self, msg))

    @staticmethod
    def split_every_five(self, encoded):
        return ' '.join([encoded[i:i + 5] for i in range(0, len(encoded), 5)])

    @staticmethod
    def encoded(self, msg):
        return ''.join(([char if char.isdigit() else self.CIPHER[char]
                         for char in self.clean(self, msg)]))

    @staticmethod
    def clean(self, msg):
        return (char for char in msg.lower() if char not in self.EXCLUDE)

    @staticmethod
    def decode(self, msg):
        return self.encoded(self, msg)


def encode(msg):
    return Atbash.encode(Atbash, msg)


def decode(msg):
    return Atbash.decode(Atbash, msg)
