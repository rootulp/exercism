def binary_search(list_of_numbers, number_to_find):
    return binary_search_with_bounds(list_of_numbers, number_to_find, 0, len(list_of_numbers) - 1)

def binary_search_with_bounds(list_of_numbers, number_to_find, left_index, right_index):
    if (left_index <= right_index):
        middle_index = left_index + (right_index - left_index) // 2
        middle = list_of_numbers[middle_index]

        if (number_to_find == middle):
            return middle_index
        elif (number_to_find < middle):
            return binary_search_with_bounds(list_of_numbers, number_to_find, left_index, middle_index - 1)
        elif (number_to_find > middle):
            return binary_search_with_bounds(list_of_numbers, number_to_find, middle_index + 1, right_index)
    else:
        raise ValueError("Value not found")
