require 'minitest/autorun'
require_relative 'garden'

class GardenTest < Minitest::Test
  def test_alices_garden
    garden = Garden.new("RC\nGG")
    assert_equal [:radishes, :clover, :grass, :grass], garden.alice
  end

  def test_different_garden_for_alice
    garden = Garden.new("VC\nRC")
    assert_equal [:violets, :clover, :radishes, :clover], garden.alice
  end

  def test_bobs_garden
    garden = Garden.new("VVCG\nVVRC")
    assert_equal [:clover, :grass, :radishes, :clover], garden.bob
  end

  def test_bob_and_charlies_gardens
    garden = Garden.new("VVCCGG\nVVCCGG")
    assert_equal [:clover, :clover, :clover, :clover], garden.bob
    assert_equal [:grass, :grass, :grass, :grass], garden.charlie
  end
end

class TestFullGarden < Minitest::Test
  def setup
    diagram = "VRCGVVRVCGGCCGVRGCVCGCGV\nVRCCCGCRRGVCGCRVVCVGCGCV"
    @garden = Garden.new(diagram)
  end

  attr_reader :garden

  def test_alice
    assert_equal [:violets, :radishes, :violets, :radishes], garden.alice
  end

  def test_bob
    assert_equal [:clover, :grass, :clover, :clover], garden.bob
  end

  def test_charlie
    assert_equal [:violets, :violets, :clover, :grass], garden.charlie
  end

  def test_david
    assert_equal [:radishes, :violets, :clover, :radishes], garden.david
  end

  def test_eve
    assert_equal [:clover, :grass, :radishes, :grass], garden.eve
  end

  def test_fred
    assert_equal [:grass, :clover, :violets, :clover], garden.fred
  end

  def test_ginny
    assert_equal [:clover, :grass, :grass, :clover], garden.ginny
  end

  def test_harriet
    assert_equal [:violets, :radishes, :radishes, :violets], garden.harriet
  end

  def test_ileana
    assert_equal [:grass, :clover, :violets, :clover], garden.ileana
  end

  def test_joseph
    assert_equal [:violets, :clover, :violets, :grass], garden.joseph
  end

  def test_kincaid
    assert_equal [:grass, :clover, :clover, :grass], garden.kincaid
  end

  def test_larry
    assert_equal [:grass, :violets, :clover, :violets], garden.larry
  end
end

class DisorderedTest < Minitest::Test
  def setup
    diagram = "VCRRGVRG\nRVGCCGCV"
    students = %w(Samantha Patricia Xander Roger)
    @garden = Garden.new(diagram, students)
  end

  attr_reader :garden

  def test_patricia
    assert_equal [:violets, :clover, :radishes, :violets], garden.patricia
  end

  def test_roger
    assert_equal [:radishes, :radishes, :grass, :clover], garden.roger
  end

  def test_samantha
    assert_equal [:grass, :violets, :clover, :grass], garden.samantha
  end

  def test_xander
    assert_equal [:radishes, :grass, :clover, :violets], garden.xander
  end
end

class TwoGardensDifferentStudents < Minitest::Test
  def diagram
    "VCRRGVRG\nRVGCCGCV"
  end

  def garden_1
    @garden_1 ||= Garden.new(diagram, %w(Alice Bob Charlie Dan))
  end

  def garden_2
    @garden_2 ||= Garden.new(diagram, %w(Bob Charlie Dan Erin))
  end

  def test_bob_and_charlie_per_garden
    assert_equal [:radishes, :radishes, :grass, :clover], garden_1.bob
    assert_equal [:violets, :clover, :radishes, :violets], garden_2.bob
    assert_equal [:grass, :violets, :clover, :grass], garden_1.charlie
    assert_equal [:radishes, :radishes, :grass, :clover], garden_2.charlie
  end
end
