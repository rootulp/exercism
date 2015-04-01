class Robot

  attr_reader :bearing, :x, :y

  DIRECTIONS = %w(north east south west).map(&:to_sym)

  def orient(direction)
    raise ArgumentError unless DIRECTIONS.include?(direction)
    @bearing = direction
  end

  def turn_right
    current = DIRECTIONS.index(bearing)
    turn(current + 1)
  end

  def turn_left
    current = DIRECTIONS.index(bearing)
    turn(current - 1)
  end

  def turn(index)
    if index > DIRECTIONS.size - 1
      @bearing = DIRECTIONS.first
    elsif index < 0
      @bearing = DIRECTIONS.last
    else
      @bearing = DIRECTIONS[index]
    end
  end

  def coordinates
    [x, y]
  end

  def at(x, y)
    @x = x
    @y = y
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

end

class Simulator

  INSTRUCTIONS = {
    'L' => :turn_left,
    'R' => :turn_right,
    'A' => :advance
  }

  def instructions(list)
    list.split(//).map {|x| INSTRUCTIONS[x]}
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
