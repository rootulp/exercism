class SpaceAge:

    SECONDS_IN_YEAR = 31557600.0
    ORBITAL_PERIODS = {
        'mercury': 0.2408467,
        'venus':   0.61519726,
        'earth':   1,
        'mars':    1.8808158,
        'jupiter': 11.862615,
        'saturn':  29.447498,
        'uranus':  84.016846,
        'neptune': 164.79132
    }

    def __init__(self, seconds):
        self.seconds = seconds

    def to_earth(self):
        return self.seconds / self.SECONDS_IN_YEAR

    def on_planet(self, planet):
        return round(self.to_earth() / self.ORBITAL_PERIODS[planet], 2)


def add_on_planet_fn(planet):
    setattr(SpaceAge, 'on_' + planet, lambda self: self.on_planet(planet))

for planet in SpaceAge.ORBITAL_PERIODS.keys():
    add_on_planet_fn(planet)
