require_relative '../../test/test_helper'
require_relative 'pascals_triangle'

class PascalsTriangleTest < Minitest::Test
  def test_one_row
    triangle = PascalsTriangle.new(1)
    assert_equal [[1]], triangle.rows
  end

  def test_two_rows
    triangle = PascalsTriangle.new(2)
    assert_equal [[1], [1, 1]], triangle.rows
  end

  def test_three_rows
    triangle = PascalsTriangle.new(3)
    assert_equal [[1], [1, 1], [1, 2, 1]], triangle.rows
  end

  def test_fourth_row
    triangle = PascalsTriangle.new(4)
    assert_equal [1, 3, 3, 1], triangle.rows.last
  end

  def test_fifth_row
    triangle = PascalsTriangle.new(5)
    assert_equal [1, 4, 6, 4, 1], triangle.rows.last
  end

  def test_twentieth_row
    triangle = PascalsTriangle.new(20)
    expected = [
      1, 19, 171, 969, 3876, 11_628, 27_132, 50_388, 75_582, 92_378, 92_378,
      75_582, 50_388, 27_132, 11_628, 3876, 969, 171, 19, 1
    ]
    assert_equal expected, triangle.rows.last
  end
end
