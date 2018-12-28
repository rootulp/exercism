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
    pass


def length(xs):
    pass


def map_clone(function, xs):
    pass


def foldl(function, xs, acc):
    pass


def foldr(function, xs, acc):
    pass


def reverse(xs):
    pass
