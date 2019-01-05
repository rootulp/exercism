def slices(string, length):
    if length <= 0 or length > len(string):
        raise ValueError('Invalid slice length')

    nums = list(map(int, list(string)))
    return [nums[i : i + length] for i in range(len(nums) - length + 1)]
