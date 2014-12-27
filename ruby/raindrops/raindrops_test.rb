require 'minitest/autorun'
require_relative 'raindrops'

class RaindropsTest < MiniTest::Unit::TestCase
  def test_1
    assert_equal '1', Raindrops.convert(1)
  end

  def test_3
    assert_equal 'Pling', Raindrops.convert(3)
  end

  def test_5
    assert_equal 'Plang', Raindrops.convert(5)
  end

  def test_7
    assert_equal 'Plong', Raindrops.convert(7)
  end

  def test_6
    assert_equal 'Pling', Raindrops.convert(6)
  end

  def test_9
    assert_equal 'Pling', Raindrops.convert(9)
  end

  def test_10
    assert_equal 'Plang', Raindrops.convert(10)
  end

  def test_14
    assert_equal 'Plong', Raindrops.convert(14)
  end

  def test_15
    assert_equal 'PlingPlang', Raindrops.convert(15)
  end

  def test_21
    assert_equal 'PlingPlong', Raindrops.convert(21)
  end

  def test_25
    assert_equal 'Plang', Raindrops.convert(25)
  end

  def test_35
    assert_equal 'PlangPlong', Raindrops.convert(35)
  end

  def test_49
    assert_equal 'Plong', Raindrops.convert(49)
  end

  def test_52
    assert_equal '52', Raindrops.convert(52)
  end

  def test_105
    assert_equal 'PlingPlangPlong', Raindrops.convert(105)
  end

  def test_12121
    assert_equal '12121', Raindrops.convert(12_121)
  end
end
