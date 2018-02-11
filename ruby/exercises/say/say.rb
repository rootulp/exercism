# Say
class Say
  UNITS = {
    1 => 'thousand',
    2 => 'million',
    3 => 'billion'
  }.freeze
  TENS = {
    90 => 'ninety',
    80 => 'eighty',
    70 => 'seventy',
    60 => 'sixty',
    50 => 'fifty',
    40 => 'forty',
    30 => 'thirty',
    20 => 'twenty'
  }.freeze
  NUMS = {
    19 => 'nineteen',
    17 => 'seventeen',
    16 => 'sixteen',
    15 => 'fifteen',
    14 => 'fourteen',
    13 => 'thirteen',
    12 => 'twelve',
    11 => 'eleven',
    10 => 'ten',
    9 => 'nine',
    8 => 'eight',
    7 => 'seven',
    6 => 'six',
    5 => 'five',
    4 => 'four',
    3 => 'three',
    2 => 'two',
    1 => 'one'
  }.freeze

  ACCEPTED_RANGE = 0...1_000_000_000_000

  attr_reader :number
  def initialize(number)
    fail ArgumentError unless ACCEPTED_RANGE.include?(number)
    @number = number
  end

  def in_english
    number_to_words(number)
  end

  private

  def number_to_words(number)
    return 'zero' if number == 0
    result = ''
    chunkify(number).each_with_index do |chunk, index|
      val = chunk_for(chunk)
      units = UNITS[index]
      result.prepend("#{val} #{units} ") if val
    end
    result.strip
  end

  def chunkify(number)
    chunks = []
    while number > 0
      number, chunk = number.divmod(1000)
      chunks << chunk
    end
    chunks
  end

  def chunk_for(number)
    return NUMS[number] if number < 20
    general_case(number)
  end

  def general_case(number)
    hundreds_digit, leftover_digits = number.divmod(100)
    tens_digit, ones_digit = leftover_digits.divmod(10)
    hundreds = NUMS[hundreds_digit] || false
    tens = TENS[tens_digit * 10] || false
    ones = NUMS[ones_digit] || false
    format(hundreds, tens, ones)
  end

  def format(hundreds, tens, ones)
    result = ''
    result += "#{hundreds} hundred " if hundreds
    result += tens.to_s if tens
    result += "-#{ones}" if ones
    result.strip
  end
end
