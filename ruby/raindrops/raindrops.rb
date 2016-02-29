require 'Prime'

# Raindrops
class Raindrops
  class << self
    WORDS = {
      3 => 'Pling',
      5 => 'Plang',
      7 => 'Plong'
    }.freeze
    def convert(num)
      return num.to_s unless valid?(factors(num))
      translate(factors(num))
    end

    def factors(num)
      Prime.prime_division(num).flatten.uniq.sort
    end

    def valid?(factors)
      factors.include?(3) || factors.include?(5) || factors.include?(7)
    end

    def translate(factors)
      factors.map { |factor| translate_factor(factor) }.join
    end

    def translate_factor(factor)
      WORDS.include?(factor) ? WORDS[factor] : ''
    end
  end
end
