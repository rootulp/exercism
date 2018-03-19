# Pythagorean Theorem
# a**2 + b**2 = c**2

# Euclid's Formula
# (m**2 - n**2) + (2mn) = (m**2 + n**2)

require 'set'

# Triplet
class Triplet
  # rubocop:disable Naming/UncommunicativeMethodParamName
  attr_reader :a, :b, :c, :sides
  def initialize(a, b, c)
    @a = a
    @b = b
    @c = c
    @sides = Set.new [a, b, c]
  end
  # rubocop:enable Naming/UncommunicativeMethodParamName

  def sum
    sides.reduce(:+)
  end

  def product
    sides.reduce(:*)
  end

  def pythagorean?
    a**2 + b**2 == c**2
  end

  def self.where(constraints = {})
    min_factor = constraints.fetch(:min_factor, 1)
    max_factor = constraints.fetch(:max_factor)
    sum = constraints.fetch(:sum, nil)

    find_triplets(min_factor, max_factor, sum)
  end

  def self.find_triplets(min_factor, max_factor, sum)
    triplets = []
    (min_factor..max_factor).to_a.combination(3).each do |a, b, c|
      triplets << Triplet.new(a, b, c) if valid?(Triplet.new(a, b, c), sum)
    end
    triplets
  end

  def self.valid?(triplet, sum)
    valid_sum?(triplet, sum) && valid_pythagorean?(triplet)
  end

  def self.valid_sum?(triplet, sum)
    return true if sum.nil?
    sum == triplet.sum
  end

  def self.valid_pythagorean?(triplet)
    triplet.pythagorean?
  end
end
