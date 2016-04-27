class TwelveDays:
    CARDINALS = {
        1: 'first',
        2: 'second',
        3: 'third',
        4: 'fourth',
        5: 'fifth',
        6: 'sixth',
        7: 'seventh',
        8: 'eighth',
        9: 'ninth',
        10: 'tenth',
        11: 'eleventh',
        12: 'twelfth'
    }

    PHRASES = {
        2: 'two Turtle Doves',
        3: 'three French Hens',
        4: 'four Calling Birds',
        5: 'five Gold Rings',
        6: 'six Geese-a-Laying',
        7: 'seven Swans-a-Swimming',
        8: 'eight Maids-a-Milking',
        9: 'nine Ladies Dancing',
        10: 'ten Lords-a-Leaping',
        11: 'eleven Pipers Piping',
        12: 'twelve Drummers Drumming'
    }

    @classmethod
    def verses(cls, start, stop):
        return "\n".join([cls.verse(i) for i in range(start, stop + 1)]) + "\n"

    @classmethod
    def verse(cls, verse_num):
        return ", ".join(filter(None, [cls.head(verse_num), cls.mid(verse_num),
                                       cls.tail(verse_num)]))

    @classmethod
    def head(cls, verse_num):
        return ("On the %(cardinality)s day of Christmas my true love gave to "
                "me" % ({"cardinality": cls.CARDINALS[verse_num]}))

    @staticmethod
    def tail(verse_num):
        if verse_num == 1:
            return "a Partridge in a Pear Tree.\n"
        return "and a Partridge in a Pear Tree.\n"

    @classmethod
    def mid(cls, verse_num):
        if verse_num != 1:
            return ", ".join([cls.PHRASES[i] for i in range(verse_num, 1, -1)])


def verse(verse_num):
    return TwelveDays.verse(verse_num)


def verses(start, stop):
    return TwelveDays.verses(start, stop)


def sing():
    return TwelveDays.verses(1, 12)
