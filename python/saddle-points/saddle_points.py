class SaddlePoints:
    def __init__(self, matrix):
        self.matrix = matrix
        self.columns = zip(*self.matrix)

    def get_saddle_points(self):
        if self.invalid_matrix():
            raise ValueError('Matrix has rows of unequal length')
        return self.saddle_points()

    def saddle_points(self):
        saddle_points = set()
        for row in range(len(self.matrix)):
            for col in range(len(self.matrix[row])):
                if self.saddle_point(row, col):
                    saddle_points.add((row, col))
        return saddle_points

    def saddle_point(self, row, col):
        return (self.matrix[row][col] == max(self.matrix[row]) and
                self.matrix[row][col] == min(self.columns[col]))

    def invalid_matrix(self):
        for row in range(len(self.matrix)):
            if len(self.matrix[row]) != len(self.matrix[0]):
                return True
        return False


def saddle_points(inp):
    return SaddlePoints(inp).get_saddle_points()
