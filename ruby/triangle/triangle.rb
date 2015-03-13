class Triangle

  attr_accessor :side_a, :side_b, :side_c
  def initialize(a, b, c)
    @side_a = a
    @side_b = b
    @side_c = c
  end

  def kind
    return :equilateral if equilateral?
    return :isosceles   if isosceles?
    return :scalene
  end

=begin
  def error?
    return true if side_a + side_b <= side_c
    return true if side_b + side_c <= side_a
    return true if side_a + side_c <= side_b
    false
  end
=end

  def equilateral?
    side_a == side_b && side_b === side_c
  end

  def isosceles?
    side_a == side_b || side_b == side_c || side_a == side_c
  end

end
