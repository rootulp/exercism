def raindrops(number):
    raindropMessage = convert_to_raindrop_message(number)
    if raindropMessage:
        return raindropMessage
    return "{}".format(number)

def convert_to_raindrop_message(number):
    raindropMessage = ""
    if is_three_a_factor(number):
        raindropMessage += "Pling"
    if is_five_a_factor(number):
        raindropMessage += "Plang"
    if is_seven_a_factor(number):
        raindropMessage += "Plong"
    return raindropMessage

def is_three_a_factor(number):
    return is_factor(number, 3)

def is_five_a_factor(number):
    return is_factor(number, 5)

def is_seven_a_factor(number):
    return is_factor(number, 7)

def is_factor(number, potentialFactor):
    return number % potentialFactor == 0

