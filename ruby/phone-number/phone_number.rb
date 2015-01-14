class PhoneNumber

  attr_reader :number, :area_code
  INVALID = '0' * 10

  def initialize(input)
    if contains_letters?(input)
      @number = INVALID
    else
      @number = clean_up(input)
    end
    @area_code = number[0..2]
  end

  def to_s
    '(' + number[0..2] + ') ' + number[3..5] + '-' + number[6..9]
  end

  private

  def clean_up(input)
    regex = /\D/
    stripped = input.gsub(regex, '')
    if stripped.length == 11 && stripped[0] == '1'
      stripped[1..-1]
    elsif stripped.length == 10
      stripped
    else
      INVALID
    end
  end

  def contains_letters?(input)
    return true if input.match(/[a-zA-Z]/)
    false
  end
    
end