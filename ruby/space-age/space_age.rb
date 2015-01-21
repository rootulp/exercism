class SpaceAge

  attr_reader :seconds

  EARTH_ORBITAL = 31557600.0
  EARTH_TO_MERCURY = 0.2408467
  EARTH_TO_VENUS = 0.61519726
  EARTH_TO_MARS = 1.8808158
  EARTH_TO_JUPITER = 11.862615
  EARTH_TO_SATURN = 29.447498
  EARTH_TO_URANUS = 84.016846
  EARTH_TO_NEPTUNE = 164.79132

  def initialize(seconds)
    @seconds = seconds
  end

  def to_earth
    seconds / EARTH_ORBITAL
  end

  def on_earth
    to_earth.round(2)
  end

  def on_mercury
    (to_earth / EARTH_TO_MERCURY).round(2)
  end

  def on_venus
    (to_earth / EARTH_TO_VENUS).round(2)
  end

  def on_mars
    (to_earth / EARTH_TO_MARS).round(2)
  end

  def on_jupiter
    (to_earth / EARTH_TO_JUPITER).round(2)
  end

  def on_saturn
    (to_earth / EARTH_TO_SATURN).round(2)
  end

  def on_uranus
    (to_earth / EARTH_TO_URANUS).round(2)
  end

  def on_neptune
    (to_earth / EARTH_TO_NEPTUNE).round(2)
  end

end