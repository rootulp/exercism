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

    def on_mercury(self):
        return self.on_planet('mercury')

    def on_venus(self):
        return self.on_planet('venus')

    def on_earth(self):
        return self.on_planet('earth')

    def on_mars(self):
        return self.on_planet('mars')

    def on_jupiter(self):
        return self.on_planet('jupiter')

    def on_saturn(self):
        return self.on_planet('saturn')

    def on_uranus(self):
        return self.on_planet('uranus')

    def on_neptune(self):
        return self.on_planet('neptune')
