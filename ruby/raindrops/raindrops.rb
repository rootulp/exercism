require 'Prime'

class Raindrops

  def self.convert(num)

    factors = self.get_factors(num)

    return num.to_s unless self.valid?(factors)
    
    self.translate(factors)
  end

  def self.get_factors(num)
    factors = Prime.prime_division(num)
    factors.flatten!
    factors.uniq!
    factors.sort!
  end

  def self.valid?(factors)
    return true if factors.include?(3) || factors.include?(5) || factors.include?(7)
    false
  end

  def self.translate(factors)
    result = ""
    factors.each do |factor|
      result << self.translate_factor(factor)
    end

    result
  end

  def self.translate_factor(factor)
    case factor
    when 3
      return "Pling"
    when 5
      return "Plang"
    when 7
      return "Plong"
    else
      return ""
    end
  end

end
