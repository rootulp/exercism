# Solution essentially taken from:
# https://github.com/xavier/exercism-assignments/blob/master/ruby/allergies/allergies.rb

class Allergies

  VALUES = {
    'eggs'         => 1,
    'peanuts'      => 2,
    'shellfish'    => 4,
    'strawberries' => 8,
    'tomatoes'     => 16,
    'chocolate'    => 32,
    'pollen'       => 64,
    'cats'         => 128
  }

  attr_reader :score
  def initialize(score)
    @score = score
  end

  def allergic_to?(item)
    VALUES[item] & score != 0
  end

  def list
    VALUES.keys.select { |item| allergic_to?(item) }
  end

end
