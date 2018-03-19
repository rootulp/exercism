# Triangle
class Triangle
  attr_accessor :side_a, :side_b, :side_c
  def initialize(side_a, side_b, side_c)
    @side_a = side_a
    @side_b = side_b
    @side_c = side_c
  end

  def kind
    raise TriangleError if error?
    return :equilateral if equilateral?
    return :isosceles   if isosceles?
    :scalene
  end

  private

  def error?
    triangle_inequality?(side_a, side_b, side_c) ||
      triangle_inequality?(side_b, side_c, side_a) ||
      triangle_inequality?(side_a, side_c, side_b)
  end

  def triangle_inequality?(side_x, side_y, side_z)
    side_x + side_y <= side_z
  end

  def equilateral?
    side_a == side_b && side_b == side_c
  end

  def isosceles?
    side_a == side_b || side_b == side_c || side_a == side_c
  end
end

TriangleError = Class.new(StandardError)
