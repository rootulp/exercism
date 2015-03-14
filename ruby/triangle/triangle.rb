class Triangle
  attr_accessor :side_a, :side_b, :side_c
  def initialize(a, b, c)
    @side_a = a
    @side_b = b
    @side_c = c
  end

  def kind
    raise TriangleError if error?
    return :equilateral if equilateral?
    return :isosceles   if isosceles?
    return :scalene
  end

  private

  def error?
    triangle_inequality?(side_a, side_b, side_c) ||
    triangle_inequality?(side_b, side_c, side_a) ||
    triangle_inequality?(side_a, side_c, side_b)
  end

  def triangle_inequality?(x, y, z)
    x + y <= z
  end

  def equilateral?
    side_a == side_b && side_b === side_c
  end

  def isosceles?
    side_a == side_b || side_b == side_c || side_a == side_c
  end
end

class TriangleError < StandardError
end
