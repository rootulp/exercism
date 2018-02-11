require 'prime'

# Prime Factors
class PrimeFactors
  def self.for(num)
    Prime.prime_division(num).map do |factors|
      Array.new(factors.last, factors.first)
    end.flatten
  end
end
