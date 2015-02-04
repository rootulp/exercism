# Influenced heavily by:
# https://github.com/xavier/exercism-assignments/blob/master/ruby/say/say.rb

class Say

  BIGS = {
    1_000_000_000 => "billion",
    1_000_000     => "million",
    1_000         => "thousand",
    100           => "hundred"
  }

  TENS = {
    90  => "ninety",
    80  => "eighty",
    70  => "seventy",
    60  => "sixty",
    50  => "fifty",
    40  => "forty",
    30  => "thirty",
    20  => "twenty"
  }

  TEENS = {
    19  => "nineteen",
    17  => "seventeen",
    16  => "sixteen",
    15  => "fifteen",
    14  => "fourteen",
    13  => "thirteen",
    12  => "twelve",
    11  => "eleven",
    10  => "ten"
  }

  DIGITS = {
    9   => "nine",
    8   => "eight",
    7   => "seven",
    6   => "six",
    5   => "five",
    4   => "four",
    3   => "three",
    2   => "two",
    1   => "one",
    0   => "zero"
  }

  ACCEPTED_RANGE = 0...1000000000000

  def initialize(number)
    raise ArgumentError unless ACCEPTED_RANGE.include?(number)
    @number = number
  end

  def in_english
    number_to_words(@number)
  end

  private

  def number_to_words(number)
    return DIGITS[number]  if number < 10
    return TEENS[number] if number < 20
    return number_to_words_100(number) if number < 100
    return number_to_words_1000(number) if number < 1000
  end

  def number_to_words_100(number)
    tens_digit, ones_digit = number.divmod(10)
    tens = TENS[tens_digit * 10]
    ones = DIGITS[ones_digit]
    ones_digit == 0 ? "#{tens}" : "#{tens}-#{ones}"
  end

  def number_to_words_1000(number)
    hundreds_digit, leftover_digits = number.divmod(100)
    hundreds = DIGITS[hundreds_digit]
    leftovers = number_to_words_100(leftover_digits)
    leftover_digits == 0 ? "#{hundreds} hundred" : "#{hundreds} hundred #{leftovers}"
  end

end