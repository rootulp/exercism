def binary_search(array, value):
    return binary_search_with_bounds(array, value, 0, len(array) - 1)

def binary_search_with_bounds(array, value, left, right):
    if (left > right):
        raise ValueError("Value not found")

    middle = left + (right - left) // 2
    middle_value = array[middle]

    if (value == middle_value):
        return middle
    elif (value < middle_value):
        return binary_search_with_bounds(array, value, left, middle - 1)
    elif (value > middle_value):
        return binary_search_with_bounds(array, value, middle + 1, right)
