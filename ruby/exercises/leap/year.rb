# Year
class Year
  def self.leap?(year)
    return true if divisible_by(year, 400)
    return false if divisible_by(year, 100)
    return true if divisible_by(year, 4)
    false
  end

  def self.divisible_by(year, divisor)
    year % divisor == 0
  end
end
