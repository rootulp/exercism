class Luhn

  attr_reader :num
  def initialize(num)
    @num = num
  end

  def addends
    digits = num.to_s.chars.map(&:to_i)
    digits.reverse
          .map.with_index {|d, i| double_helper(d, i)}
          .reverse
  end

  def checksum
    addends.inject(:+)
  end

  def valid?
    checksum % 10 == 0
  end

  def self.create(num)
    (0..9).each do |i|
      temp = Luhn.new("#{num}#{i}".to_i)
      return temp.num if temp.valid?
    end
  end

  private

  def double_helper(digit, index)
    digit *= 2 if index % 2 == 1
    digit -= 9 if digit > 9
    digit
  end

end
