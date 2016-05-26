class Triangle:

    EQUILATERAL = "equilateral"
    ISOSCELES = "isosceles"
    SCALENE = "scalene"

    def __init__(self, a, b, c):
        self.a = a
        self.b = b
        self.c = c
        if self.error():
            raise TriangleError

    def kind(self):
        if self.equilateral():
            return self.EQUILATERAL
        if self.isosceles():
            return self.ISOSCELES
        return self.SCALENE

    def equilateral(self):
        return self.a == self.b == self.c

    def isosceles(self):
        return self.a == self.b or self.b == self.c or self.a == self.c

    def error(self):
        return self.negative_sides() or self.triangle_inequality()

    def negative_sides(self):
        return self.a <= 0 or self.b <= 0 or self.c <= 0

    def triangle_inequality(self):
        return (self.a + self.b <= self.c or
                self.b + self.c <= self.a or
                self.a + self.c <= self.b)


class TriangleError(Exception):
    pass
