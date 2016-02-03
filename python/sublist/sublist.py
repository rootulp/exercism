SUBLIST = 0
SUPERLIST = 1
EQUAL = 2
UNEQUAL = 3


def check_lists(l1, l2):
    if l1 == l2:
        return EQUAL
    elif l1 in each_cons(l2, len(l1)):
        return SUBLIST
    elif l2 in each_cons(l1, len(l2)):
        return SUPERLIST
    else:
        return UNEQUAL


# Not the most efficent
def each_cons(lst, size):
    return [lst[i: i + size] for i in range(len(lst) - size + 1)]
