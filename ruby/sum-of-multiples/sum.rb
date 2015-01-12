# Probably a better way to do this then having
# seperate class and instance methods

class SumOfMultiples

  @@multiples = [3,5]
  
  def self.to(max)
    sum = 0
    (1..max-1).step() do |curr|
      sum += curr if self.divisible?(curr)
    end
    sum
  end

  def self.divisible?(num)
    @@multiples.each do |multiple|
      return true if num % multiple == 0
    end
    false
  end

  def initialize(*multiples)
    @multiples = multiples
  end

  def to(max)
    sum = 0
    (1..max-1).step() do |curr|
      sum += curr if self.divisible?(curr)
    end
    sum
  end

  def divisible?(num)
    @multiples.each do |multiple|
      return true if num % multiple == 0
    end
    false
  end

end
