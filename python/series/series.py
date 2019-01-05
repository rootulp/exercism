def slices(string, size):
    if size <= 0 or size > len(string):
        raise ValueError('Invalid slice size')

    return [string[i:i+size] for i in range(len(string) - size + 1)]
