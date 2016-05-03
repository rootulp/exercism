require_relative '../../test/test_helper'
require_relative 'triplet'

class TripletTest < Minitest::Test
  def test_sum
    assert_equal 12, Triplet.new(3, 4, 5).sum
  end

  def test_product
    assert_equal 60, Triplet.new(3, 4, 5).product
  end

  def test_pythagorean
    assert Triplet.new(3, 4, 5).pythagorean?
  end

  def test_not_pythagorean
    assert !Triplet.new(5, 6, 7).pythagorean?
  end

  def test_triplets_upto_10
    triplets = Triplet.where(max_factor: 10)
    products = triplets.map(&:product).sort
    assert_equal [60, 480], products
  end

  def test_triplets_from_11_upto_20
    triplets = Triplet.where(min_factor: 11, max_factor: 20)
    products = triplets.map(&:product).sort
    assert_equal [3840], products
  end

  def test_triplets_where_sum_x
    triplets = Triplet.where(sum: 180, max_factor: 100)
    products = triplets.map(&:product).sort
    assert_equal [118_080, 168_480, 202_500], products
  end
end
