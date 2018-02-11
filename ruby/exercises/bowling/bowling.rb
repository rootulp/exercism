# Game
class Game
  attr_accessor :frames, :penultimate_roll, :last_roll
  def initialize
    @frames = []
    @penultimate_roll = 0
    @last_roll = 0
  end

  def roll(pins)
    raise RuntimeError if invalid_pins?(pins)
    if frames.any? && frames.last.second_throw.nil?
      frames.last.second_throw = pins
    else
      frames << Frame.new(pins)
    end
  end

  def score
    frames.reduce(0) do |sum, frame|
      sum + frame.score
    end
  end

  def invalid_pins?(pins)
    pins < 0 || pins > 10
  end
end

class Frame
  attr_accessor :first_throw, :second_throw
  def initialize(first_throw, second_throw = nil, multiplier = 1)
    @first_throw = first_throw
    @second_throw = second_throw
    @second_throw = 0 if strike?(first_throw)
    @multiplier = multiplier
  end

  def score
    first_throw + second_throw
  end

  def strike?(pins)
    pins == 10
  end
end
