def slices(string, length):
    if length <= 0 or length > len(string):
        raise ValueError('Invalid slice length')

    return [string[i:i+length] for i in range(len(string) - length + 1)]
