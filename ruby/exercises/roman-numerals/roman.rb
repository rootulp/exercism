# Fixnum
class Integer # I don't like monkey patching
  NUMERALS = {
    1 => 'I',
    4 => 'IV',
    5 => 'V',
    9 => 'IX',
    10 => 'X',
    40 => 'XL',
    50 => 'L',
    90 => 'XC',
    100 => 'C',
    400 => 'CD',
    500 => 'D',
    900 => 'CM',
    1000 => 'M'
  }.freeze

  def to_roman
    result = ''
    target = self

    NUMERALS.keys.reverse_each do |num|
      while target >= num
        target -= num
        result += NUMERALS[num]
      end
    end
    result
  end
end
