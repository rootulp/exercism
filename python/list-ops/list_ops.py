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
    pass


def foldl(function, xs, acc):
    pass


def foldr(function, xs, acc):
    pass


def reverse(xs):
    pass
