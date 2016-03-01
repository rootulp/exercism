# Robot
class Robot
  BEARINGS = [:north, :east, :south, :west].freeze
  attr_reader :bearing, :x, :y

  def orient(bearing)
    raise ArgumentError unless BEARINGS.include?(bearing)
    @bearing = bearing
  end

  def at(x, y)
    @x = x
    @y = y
  end

  def turn_right
    turn(bearing_index + 1)
  end

  def turn_left
    turn(bearing_index - 1)
  end

  def advance
    if bearing == :north
      @y += 1
    elsif bearing == :east
      @x += 1
    elsif bearing == :south
      @y -= 1
    elsif bearing == :west
      @x -= 1
    end
  end

  def coordinates
    [x, y]
  end

  private

  def turn(index)
    @bearing = wrap_turn(index)
  end

  def bearing_index
    BEARINGS.index(bearing)
  end

  def wrap_turn(index)
    if index > BEARINGS.size - 1
      BEARINGS.first
    elsif index < 0
      BEARINGS.last
    else
      BEARINGS[index]
    end
  end
end

class Simulator
  INSTRUCTIONS = {
    'L' => :turn_left,
    'R' => :turn_right,
    'A' => :advance
  }.freeze

  def instructions(list)
    list.chars.map { |instruction| INSTRUCTIONS[instruction] }
  end

  def place(robot, x: 0, y: 0, direction: :north)
    robot.orient(direction)
    robot.at(x, y)
  end

  def evaluate(robot, list)
    instructions(list).each do |instruction|
      robot.send(instruction)
    end
  end
end
