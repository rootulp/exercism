class SpaceAge:

    ORBITAL_PERIODS = {
        'mercury': 7600530.24,
        'venus': 19413907.2,
        'earth': 31558149.76,
        'mars': 59354294.4,
        'jupiter': 374335776.0,
        'saturn': 929596608.0,
        'uranus': 2661041808.0,
        'neptune': 5200418592.0
    }

    def __init__(self, seconds):
        self.seconds = seconds

    def on_planet(self, planet):
        return round(self.seconds / self.ORBITAL_PERIODS[planet], 2)


def add_on_planet_fn(planet):
    setattr(SpaceAge, 'on_' + planet, lambda self: self.on_planet(planet))


for planet in SpaceAge.ORBITAL_PERIODS.keys():
    add_on_planet_fn(planet)
