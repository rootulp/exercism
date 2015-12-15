class SpaceAge

  SECONDS_IN_YEAR = 31557600.0
  ORBITAL_PERIODS = {
    'mercury' => 0.2408467,
    'venus' => 0.61519726,
    'earth' => 1,
    'mars' => 1.8808158,
    'jupiter' => 11.862615,
    'saturn' => 29.447498,
    'uranus' => 84.016846,
    'neptune' => 164.79132
  }

  attr_reader :seconds
  def initialize(seconds)
    @seconds = seconds
  end

  def to_earth
    seconds / SECONDS_IN_YEAR
  end

  def on_planet(planet)
    (to_earth / ORBITAL_PERIODS[planet]).round(2)
  end

  ORBITAL_PERIODS.keys.each do |planet|
    define_method :"on_#{planet}" do
      on_planet(planet)
    end
  end

end
