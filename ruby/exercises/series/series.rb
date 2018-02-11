# Series
class Series
  attr_reader :digits
  def initialize(digits)
    @digits = digits.chars.map(& :to_i)
  end

  def slices(length)
    fail ArgumentError if length > digits.length
    digits.each_cons(length).to_a
  end
end
