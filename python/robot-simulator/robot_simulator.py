class NORTH:
    @staticmethod
    def advance(self, x, y):
        return (x, y + 1)

    @staticmethod
    def turn_right(self):
        return EAST

    @staticmethod
    def turn_left(self):
        return WEST


class EAST:
    @staticmethod
    def advance(self, x, y):
        return (x + 1, y)

    @staticmethod
    def turn_right(self):
        return SOUTH

    @staticmethod
    def turn_left(self):
        return NORTH


class SOUTH:
    @staticmethod
    def advance(self, x, y):
        return (x, y - 1)

    @staticmethod
    def turn_right(self):
        return WEST

    @staticmethod
    def turn_left(self):
        return EAST


class WEST:
    @staticmethod
    def advance(self, x, y):
        return (x - 1, y)

    @staticmethod
    def turn_right(self):
        return NORTH

    @staticmethod
    def turn_left(self):
        return SOUTH


class Robot:

    def __init__(self, direction=NORTH, x=0, y=0):
        self.coordinates = (x, y)
        self.bearing = direction

    def advance(self):
        self.coordinates = self.bearing.advance(self.bearing,
                                                self.x(), self.y())

    def turn_right(self):
        self.bearing = self.bearing.turn_right(self.bearing)

    def turn_left(self):
        self.bearing = self.bearing.turn_left(self.bearing)

    def simulate(self, instructions):
        for i in instructions:
            self.execute_instruction(i)

    def execute_instruction(self, i):
        if i == 'A':
            self.advance()
        elif i == 'L':
            self.turn_left()
        elif i == 'R':
            self.turn_right()

    def x(self):
        return self.coordinates[0]

    def y(self):
        return self.coordinates[1]
