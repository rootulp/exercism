class Sieve

  def initialize(max)
    @range = 2..max
  end

  def primes
    primes = []
    composites = []

    @range.each do |curr|
      next if composites.include?(curr)
      primes << curr
      composites = composites | multiples(curr)
    end
    primes
  end

  def multiples(num)
    @range.select { |multiple| multiple % num == 0 }
  end

end