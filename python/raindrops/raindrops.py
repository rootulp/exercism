def raindrops(number):
    if is_three_a_factor(number):
        return "Pling"
    return "{}".format(number)

def is_three_a_factor(number):
    return number % 3 == 0
