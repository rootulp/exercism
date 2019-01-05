def slices(string, length):
    if length <= 0 or length > len(string):
        raise ValueError('Invalid slice length')

    slices = []
    for indx, _elem in enumerate(string):
        if len(string) - indx >= length:
            curr_slice = []
            curr_indx = indx
            while len(curr_slice) < length:
                curr_slice.append(int(string[curr_indx]))
                curr_indx += 1
            slices.append(curr_slice)
    return slices
