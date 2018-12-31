class Matrix:

    def __init__(self, inp):
        self.rows = self.build_rows(inp)
        self.columns = list(map(list, list(zip(*self.rows))))

    @staticmethod
    def build_rows(inp):
        return [[int(val) for val in row.split()] for row in inp.split('\n')]
