class Triplet

  # Pythagorean Theorem
  # a**2 + b**2 = c**2

  # Euclid's Formula
  # (m**2 - n**2) + (2mn) = (m**2 + n**2)

  attr_reader :a, :b, :c
  def initialize(a, b, c)
    @a = a
    @b = b
    @c = c
  end

  def sum
    a + b + c
  end

  def product
    a * b * c
  end

  def pythagorean?
    a**2 + b**2 == c**2
  end

  class << self

    def where(constraints = {})
      max_factor = constraints.fetch(:max_factor)
      min_factor = constraints.fetch(:min_factor, 1)
      sum        = constraints.fetch(:sum, false)

      triplets = []

      factors = min_factor..max_factor
      factors.to_a.combination(3) do |a, b, c|
        temp = Triplet.new(a, b, c)
        triplets << temp if valid?(temp, sum)
      end

      triplets
    end

    def valid?(triplet, sum)
      if triplet.pythagorean? && sum
        triplet.sum == sum
      elsif triplet.pythagorean?
        true
      else
        false
      end
    end

  end

end
