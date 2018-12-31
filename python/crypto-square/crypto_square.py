import string
import math
import itertools


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

    @classmethod
    def transpose_square(cls, square):
        matrix = [list(row) for row in square]
        transposed_matrix = cls.filter_out_none(cls.transpose_uneven_matrix(matrix))
        return [''.join(row) for row in transposed_matrix]

    @staticmethod
    def normalize(msg):
        return ''.join(ch.lower() for ch in msg if ch not in
                       set(string.punctuation + ' '))

    @staticmethod
    def square_size(msg_length):
        return int(math.ceil(msg_length ** 0.5))

    # https://stackoverflow.com/a/4938130/2813210
    @staticmethod
    def transpose_uneven_matrix(matrix):
        return list(itertools.zip_longest(*matrix))

    @staticmethod
    def filter_out_none(matrix):
        return [[val for val in row if val is not None] for row in matrix]

def encode(msg):
    return CryptoSquare.encode(msg)
