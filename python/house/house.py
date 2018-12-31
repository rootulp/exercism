class House:

    LYRICS = [("built", "house that Jack"),
              ("ate", "malt"),
              ("killed", "rat"),
              ("worried", "cat"),
              ("tossed", "dog"),
              ("milked", "cow with the crumpled horn"),
              ("kissed", "maiden all forlorn"),
              ("married", "man all tattered and torn"),
              ("woke", "priest all shaven and shorn"),
              ("kept", "rooster that crowed in the morn"),
              ("belonged to", "farmer sowing his corn"),
              ("", "horse and the hound and the horn")]

    LAST_LINE = "that lay in the house that Jack built."

    @classmethod
    def rhyme(cls):
        return "\n\n".join([cls.verse(i) for i in range(12)])

    @classmethod
    def verse(cls, verse_num):
        return "\n".join([_f for _f in cls.parts(verse_num) if _f])

    @classmethod
    def parts(cls, verse_num):
        return [cls.first(verse_num), cls.middle(verse_num),
                cls.last(verse_num)]

    @classmethod
    def first(cls, verse_num):
        if verse_num != 0:
            return cls.first_partial(verse_num)
        return cls.first_partial(verse_num) + " " + cls.verb(verse_num) + "."

    @classmethod
    def first_partial(cls, verse_num):
        return "This is the " + cls.noun(verse_num)

    @classmethod
    def middle(cls, verse_num):
        if verse_num >= 2:
            return "\n".join([cls.middle_partial(num) for num in
                              range(verse_num - 1, 0, -1)])

    @classmethod
    def middle_partial(cls, num):
        return "that " + cls.verb(num) + " the " + cls.noun(num)

    @classmethod
    def verb(cls, num):
        return cls.LYRICS[num][0]

    @classmethod
    def noun(cls, num):
        return cls.LYRICS[num][1]

    @classmethod
    def last(cls, verse_num):
        if verse_num != 0:
            return cls.LAST_LINE


def verse(verse_num):
    return House.verse(verse_num)


def rhyme():
    return House.rhyme()
