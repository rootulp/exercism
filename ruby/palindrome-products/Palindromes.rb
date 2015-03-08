class Palindromes
  attr_accessor :max, :min, :factors
  
  DEFAULTS = {min_factor: 1}
  def initialize(args)
    @max = args[:max_factor]
    @min = args[:min_factor]
  end

  def generate
    (min..max).each do |x|
      (min..max).each do |y|
         factors << [x, y] if factors?(x, y)
      end
    end
  end

end
