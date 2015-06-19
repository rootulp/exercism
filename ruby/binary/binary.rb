class Binary

  def initialize(binary)
    @binary = binary
  end

  def to_decimal
    return 0 unless valid?

    total = 0
    @binary.split(//).reverse.each_with_index do |one_or_zero, index|
      total += 2 ** index if one_or_zero == "1"
    end
    total
  end

  def valid?
    regex = /[^01]/
    return false if @binary.match(regex)
    true
  end

end
