def raindrops(number):
    message = convert_to_message(number)
    return message if message else "{}".format(number)

def convert_to_message(number):
    return f"{'Pling' if is_three_a_factor(number) else ''}" \
           f"{'Plang' if is_five_a_factor(number) else ''}" \
           f"{'Plong' if is_seven_a_factor(number) else ''}"

def is_three_a_factor(number):
    return is_factor(number, 3)

def is_five_a_factor(number):
    return is_factor(number, 5)

def is_seven_a_factor(number):
    return is_factor(number, 7)

def is_factor(number, potentialFactor):
    return number % potentialFactor == 0
