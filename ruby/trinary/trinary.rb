class Trinary

  def initialize(trinary)
    @trinary = trinary
  end

  def to_decimal
    total = 0
    return total unless valid?

    @trinary.split(//).reverse.each_with_index do |multiplier, power|
      total += (3 ** power) * multiplier.to_i
    end

    total
  end

  def valid?
    regex = /[^0123]/
    return false if @trinary.match(regex)
    true
  end

end