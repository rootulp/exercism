# Phone Number
class PhoneNumber
  INVALID = ('0' * 10).freeze

  attr_reader :input
  def initialize(input)
    @input = input
  end

  def number
    return INVALID unless valid?
    return stripped if valid_ten_digits?
    stripped[1..-1] if valid_eleven_digits?
  end

  def area_code
    number[0..2]
  end

  def to_s
    '(' + number[0..2] + ') ' + number[3..5] + '-' + number[6..9]
  end

  private

  def valid?
    no_letters? && (valid_ten_digits? || valid_eleven_digits?)
  end

  def valid_ten_digits?
    stripped.size == 10
  end

  def valid_eleven_digits?
    stripped.size == 11 && stripped.chars.first == '1'
  end

  def stripped
    input.gsub(/\D/, '')
  end

  def no_letters?
    !input.match(/[a-zA-Z]/)
  end
end
