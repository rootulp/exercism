def slices(string, size):
    if size <= 0 or size > len(string):
        raise ValueError('Invalid slice size')

    slices = []
    for indx, elem in enumerate(string):
        if len(string) - indx >= size:
            curr_slice = []
            curr_indx = indx
            while len(curr_slice) < size:
                curr_slice.append(int(string[curr_indx]))
                curr_indx += 1
            slices.append(curr_slice)
    return slices