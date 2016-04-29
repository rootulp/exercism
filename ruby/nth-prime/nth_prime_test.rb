require 'minitest/autorun'
require_relative 'nth_prime'

class TestPrimes < Minitest::Test
  def test_first
    assert_equal 2, NthPrime.nth(1)
  end

  def test_second
    assert_equal 3, NthPrime.nth(2)
  end

  def test_sixth_prime
    assert_equal 13, NthPrime.nth(6)
  end

  def test_big_prime
    assert_equal 104_743, NthPrime.nth(10_001)
  end

  def test_weird_case
    assert_raises ArgumentError do
      NthPrime.nth(0)
    end
  end
end
