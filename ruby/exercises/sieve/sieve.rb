# Sieve
class Sieve
  DEFAULT_MIN = 2
  attr_reader :min, :max, :range
  def initialize(max)
    @min = DEFAULT_MIN
    @max = max
    @range = min..max
  end

  def primes
    primes = []
    unmarked = range.to_a

    while unmarked.any?
      prime = unmarked.shift
      unmarked -= multiples(prime)
      primes << prime
    end
    primes
  end

  def multiples(num)
    range.select { |multiple| multiple % num == 0 }
  end
end
