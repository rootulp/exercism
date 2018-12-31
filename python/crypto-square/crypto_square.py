import string
import math


class CryptoSquare:

    @classmethod
    def encode(cls, msg):
        if len(cls.normalize(msg)) == 0:
            return ''
        return ' '.join(cls.transpose_square(cls.squarify(cls.normalize(msg))))

    @classmethod
    def squarify(cls, msg):
        return [msg[i:i + cls.square_size(len(msg))]
                for i in range(0, len(msg), cls.square_size(len(msg)))]

    @staticmethod
    def transpose_square(square):
        return list(map(lambda *row: ''.join([ch if ch else '' for ch in row]),
                   *square))

    @staticmethod
    def normalize(msg):
        return ''.join(ch.lower() for ch in msg if ch not in
                       set(string.punctuation + ' '))

    @staticmethod
    def square_size(msg_length):
        return int(math.ceil(msg_length**.5))


def encode(msg):
    return CryptoSquare.encode(msg)
