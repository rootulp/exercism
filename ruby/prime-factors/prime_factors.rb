require 'prime'

class PrimeFactors

  def self.for(num)
    result = []

    Prime.prime_division(num).each do |factors|
      factors[1].times { result << factors[0] }
    end

    result
  end

end