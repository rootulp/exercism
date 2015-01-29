class Hexadecimal

  HEX_VALUES = {"0" => 0, "1" => 1, "2" => 2, "3" => 3, "4" => 4, 
                "5" => 5, "6" => 6, "7" => 7, "8" => 8, "9" => 9, 
                "a" => 10, "b" => 11, "c" => 12, 
                "d" => 13, "e" => 14, "f" => 15}

  attr_reader :hexadecimal
  def initialize(hexadecimal)
    @hexadecimal = hexadecimal
  end

  def to_decimal
    total = 0
    return total unless valid?

    hexadecimal.split(//).reverse.each_with_index do |multiplier, power|
      total += (16 ** power) * HEX_VALUES[multiplier]
    end

    total
  end

  private 

  def valid?
    hexadecimal.match(/[^a-f\d]/) ? false : true
  end

end