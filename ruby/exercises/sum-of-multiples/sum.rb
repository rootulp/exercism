# Probably a better way to do this then having
# seperate class and instance methods
class SumOfMultiples
  MULTIPLES = [3, 5].freeze
  attr_reader :multiples

  def self.to(max)
    SumOfMultiples.new.to(max)
  end

  def initialize(*multiples)
    @multiples = multiples.any? ? multiples : MULTIPLES
  end

  def to(max)
    (1...max).step.select { |num| divisible?(num) }.reduce(0, :+)
  end

  def divisible?(num)
    multiples.any? { |multiple| (num % multiple).zero? }
  end
end
