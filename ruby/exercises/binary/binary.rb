# Used to convert a binary string to an int
class Binary
  attr_reader :binary
  def initialize(binary)
    @binary = binary
  end

  def to_decimal
    return 0 if invalid?
    convert_to_decimal
  end

  private

  def invalid?
    binary.match(/[^01]/)
  end

  def convert_to_decimal
    binary.chars.reverse_each.with_index.map do |x, i|
      x == '0' ? 0 : 2**i
    end.reduce(:+)
  end
end
