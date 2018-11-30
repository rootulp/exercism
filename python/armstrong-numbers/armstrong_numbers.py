def get_digits(number):
    return list(map(int, list(str(number))))


def is_armstrong(number):
    digits = get_digits(number)
    num_digits = len(digits)
    digits_raised_to_num_digits = [digit ** num_digits for digit in digits]

    return number == sum(digits_raised_to_num_digits)
