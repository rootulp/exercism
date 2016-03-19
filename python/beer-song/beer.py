class Beer:

    LAST_LINE = ('Go to the store and buy some more, '
                 '99 bottles of beer on the wall.')

    @classmethod
    def song(cls, start, stop):
        return "\n".join([cls.verse(verse_num) for verse_num
                          in reversed(range(stop, start + 1))]) + "\n"

    @classmethod
    def verse(cls, verse_num):
        return "\n".join((cls.prefix(verse_num), cls.suffix(verse_num))) + "\n"

    @classmethod
    def prefix(cls, verse_num):
        return ('%(quantity)s %(container)s of beer on the wall, '
                '%(quantity)s %(container)s of beer.'
                % cls.vals_for(verse_num)).capitalize()

    @classmethod
    def suffix(cls, verse_num):
        if verse_num == 0:
            return cls.LAST_LINE
        else:
            return ('Take %(cardinality)s down and pass it around, '
                    '%(quantity)s %(container)s of beer on the wall.'
                    % cls.vals_for(verse_num - 1))

    @classmethod
    def vals_for(cls, num):
        return {'quantity': cls.quantity(num),
                'container': cls.container(num),
                'cardinality': cls.cardinality(num)}

    @staticmethod
    def quantity(num):
        return 'no more' if num == 0 else str(num)

    @staticmethod
    def container(num):
        return 'bottle' if num == 1 else 'bottles'

    @staticmethod
    def cardinality(num):
        return 'it' if num == 0 else 'one'


def verse(verse_num):
    return Beer.verse(verse_num)


def song(start, stop=0):
    return Beer.song(start, stop)
