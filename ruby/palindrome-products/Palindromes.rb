require 'set'

class Palindromes
  attr_accessor :min, :max, :factors

  DEFAULT_MIN = 1
  def initialize(args)
    @min = args[:min_factor] || DEFAULT_MIN
    @max = args[:max_factor]
    @factors = Set.new
  end

  def generate
    (min..max).each do |x|
      (min..max).each do |y|
        if factors?(x, y)
          if x <= y
            factors << [x, y]
          else
            factors << [y, x]
          end
        end
      end
    end
  end

  def factors?(x, y)
    result = (x * y).to_s
    result == result.reverse
  end

  def largest
    build_subset(find_largest)
  end

  def smallest
    build_subset(find_smallest)
  end

  def find_largest
    factor = factors.max_by {|x| x.reduce(:*)}
    factor.reduce(:*)
  end

  def find_smallest
    factor = factors.min_by {|x| x.reduce(:*)}
    factor.reduce(:*)
  end

  def build_subset(val)
    Subset.new(val, find_subset_factors(val))
  end

  def find_subset_factors(val)
    subset_factors = []
    factors.each do |factor|
      subset_factors << factor if factor[0] * factor[1] == val
    end
    subset_factors
  end

end

class Subset
  attr_reader :value, :factors
  def initialize(value, factors)
    @value = value
    @factors = factors
  end
end
