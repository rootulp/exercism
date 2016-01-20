class Luhn

  def initialize(num)
    @num = num
  end

  def addends
    @num.to_s.chars.map(&:to_i) # digits
        .reverse
        .map.with_index { |digit, index| addend(digit, index) }
        .reverse
  end

  def addend(digit, index)
    subtract_nine(double(digit, index))
  end

  def checksum
    addends.inject(:+)
  end

  def valid?
    checksum % 10 == 0
  end

  def self.create(num)
    (0..9).each do |i|
      return "#{num}#{i}".to_i if Luhn.new("#{num}#{i}".to_i).valid?
    end
  end

  private

  def double(digit, index)
    index % 2 == 1 ? digit * 2 : digit
  end

  def subtract_nine(digit)
    digit > 9 ? digit - 9 : digit
  end

end
