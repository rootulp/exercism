def slices(string, length):
    if length <= 0 or length > len(string):
        raise ValueError('Invalid slice length')

    nums = list(map(int, list(string)))
    return get_slices(nums, length)


def get_slices(arr, length):
    return [arr[i:i + length] for i in range(len(arr) - length + 1)]
