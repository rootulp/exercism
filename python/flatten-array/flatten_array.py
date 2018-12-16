def flatten(nested_list):
    return [element for element in flatten_list(nested_list)]


def flatten_list(nested_list):
    for element in nested_list:
        if not isinstance(element, list):
            if element is not None and not isinstance(element, tuple):
                yield element
        else:
            for subelement in flatten_list(element):
                yield subelement
