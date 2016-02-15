require 'set'

# Palindromes
class Palindromes
  DEFAULT_MIN = 1

  attr_accessor :min, :max, :factors
  def initialize(args)
    @min = args[:min_factor] || DEFAULT_MIN
    @max = args[:max_factor]
    @factors = Set.new
  end

  def generate
    (min..max).each do |x|
      (min..max).each do |y|
        next unless factors?(x, y)
        factors << (x <= y ? [x, y] : [y, x])
      end
    end
  end

  def largest
    build_subset(find_largest)
  end

  def smallest
    build_subset(find_smallest)
  end

  private

  def factors?(x, y)
    result = (x * y).to_s
    result == result.reverse
  end

  def find_largest
    factor = factors.max_by { |x| x.reduce(:*) }
    factor.reduce(:*)
  end

  def find_smallest
    factor = factors.min_by { |x| x.reduce(:*) }
    factor.reduce(:*)
  end

  def find_subset_factors(val)
    factors.select { |x| x.reduce(:*) == val }
  end

  def build_subset(val)
    Subset.new(val, find_subset_factors(val))
  end
end

# Subset
class Subset
  attr_reader :value, :factors
  def initialize(value, factors)
    @value = value
    @factors = factors
  end
end
