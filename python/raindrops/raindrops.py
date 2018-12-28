def raindrops(number):
    raindropMessage = ""
    if is_three_a_factor(number):
        raindropMessage += "Pling"
    if is_five_a_factor(number):
        raindropMessage += "Plang"
    if is_seven_a_factor(number):
        raindropMessage += "Plong"
    if raindropMessage:
        return raindropMessage
    return "{}".format(number)

def is_three_a_factor(number):
    return number % 3 == 0

def is_five_a_factor(number):
    return number % 5 == 0

def is_seven_a_factor(number):
    return number % 7 == 0
