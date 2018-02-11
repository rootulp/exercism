require_relative 'say'

class SayTest < Minitest::Test
  def test_0
    assert_equal 'zero', Say.new(0).in_english
  end

  def test_one
    assert_equal 'one', Say.new(1).in_english
  end

  def test_14
    assert_equal 'fourteen', Say.new(14).in_english
  end

  def test_twenty
    # This really shouldn't be twenty-zero
    assert_equal 'twenty', Say.new(20).in_english
  end

  def test_twenty_two
    assert_equal 'twenty-two', Say.new(22).in_english
  end

  def test_100
    assert_equal 'one hundred', Say.new(100).in_english
  end

  def test_120
    assert_equal 'one hundred twenty', Say.new(120).in_english
  end

  def test_123
    assert_equal 'one hundred twenty-three', Say.new(123).in_english
  end

  def test_1_thousand
    assert_equal 'one thousand', Say.new(1000).in_english
  end

  def test_1_thousand_234
    expected = 'one thousand two hundred thirty-four'
    assert_equal expected, Say.new(1234).in_english
  end

  def test_1_million
    assert_equal 'one million', Say.new(10**6).in_english
  end

  def test_1_million_and_some_crumbs
    assert_equal 'one million two', Say.new(1_000_002).in_english
  end

  def test_1_million_2_thousand_345
    expected = 'one million two thousand three hundred forty-five'
    assert_equal expected, Say.new(1_002_345).in_english
  end

  def test_1_billion
    assert_equal 'one billion', Say.new(10**9).in_english
  end

  def test_really_big_number
    expected = 'nine hundred eighty-seven billion '
    expected << 'six hundred fifty-four million '
    expected << 'three hundred twenty-one thousand '
    expected << 'one hundred twenty-three'
    assert_equal expected, Say.new(987_654_321_123).in_english
  end

  def test_lower_bound
    assert_raises ArgumentError do
      Say.new(-1).in_english
    end
  end

  def test_upper_bound
    assert_raises ArgumentError do
      Say.new(1_000_000_000_000).in_english
    end
  end
end
