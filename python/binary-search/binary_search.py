def binary_search(array, num_to_find):
    return binary_search_with_bounds(array, num_to_find, 0, len(array) - 1)

def binary_search_with_bounds(array, num_to_find, left, right):
    if (left <= right):
        middle = left + (right - left) // 2
        middle_value = array[middle]

        if (num_to_find == middle_value):
            return middle
        elif (num_to_find < middle_value):
            return binary_search_with_bounds(array, num_to_find, left, middle - 1)
        elif (num_to_find > middle_value):
            return binary_search_with_bounds(array, num_to_find, middle + 1, right)
    else:
        raise ValueError("Value not found")
