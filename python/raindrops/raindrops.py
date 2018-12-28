def raindrops(number):
    if is_three_a_factor(number):
        return "Pling"
    if is_five_a_factor(number):
        return "Plang"
    if is_seven_a_factor(number):
        return "Plong"
    return "{}".format(number)

def is_three_a_factor(number):
    return number % 3 == 0

def is_five_a_factor(number):
    return number % 5 == 0

def is_seven_a_factor(number):
    return number % 7 == 0
