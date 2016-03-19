def song(start, stop=0):
    return "\n".join([verse(verse_num) for verse_num
                      in reversed(range(stop, start + 1))]) + "\n"


def verse(verse_num):
    return "\n".join((prefix(verse_num), suffix(verse_num))) + "\n"


def prefix(verse_num):
    return ("%(quantity)s %(container)s of beer on the wall, %(quantity)s "
            "%(container)s of beer." % vals_for(verse_num)).capitalize()


def suffix(verse_num):
    if verse_num == 0:
        return last_line()
    else:
        return "Take %(cardinality)s down and pass it around, %(quantity)s "\
               "%(container)s of beer on the wall." % vals_for(verse_num - 1)


def vals_for(num):
    return {"quantity": quantity(num),
            "container": container(num),
            "cardinality": cardinality(num)}


def quantity(num):
    return 'no more' if num == 0 else str(num)


def container(num):
    return 'bottle' if num == 1 else 'bottles'


def cardinality(num):
    return 'it' if num == 0 else 'one'


def last_line():
    return "Go to the store and buy some more, "\
            "99 bottles of beer on the wall."
