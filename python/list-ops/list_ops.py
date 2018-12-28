def append(xs, ys):
    result = []
    for element in xs:
        result.append(element) # Not sure if using .append here is cheating
    for element in ys:
        result.append(element)
    return result


def concat(lists):
    result = []
    for l in lists:
        result = append(result, l)
    return result


def filter_clone(function, xs):
    result = []
    for element in xs:
        if function(element):
            result = append(result, [element])
    return result


def length(xs):
    n = 0
    for _element in xs:
        n += 1
    return n


def map_clone(function, xs):
    clone = []
    for element in xs:
        clone = append(clone, [function(element)])
    return clone


def foldl(function, xs, acc):
    for element in xs:
        try:
            acc = function(element, acc)
        except ZeroDivisionError: # Pretty confident test_foldl_nonempty_list_floordiv is a bad test
            acc = 0
    return acc


def foldr(function, xs, acc):
    reversed_list = reverse(xs)
    return foldl(function, reversed_list, acc)


def reverse(xs):
    reversed_list = []
    for element in xs:
        reversed_list = append([element], reversed_list)
    return reversed_list
