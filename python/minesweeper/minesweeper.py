class Minesweeper:

    MINE = "*"
    SPACE = " "
    CORNER = "+"
    VERTICAL_EDGE = "|"
    HORIZONTAL_EDGE = "-"

    @classmethod
    def board(cls, inp):
        if not cls.valid(inp):
            raise ValueError
        return cls.generate_board(inp)

    @classmethod
    def generate_board(cls, inp):
        inp = map(lambda row: list(row), inp)
        for y, row in enumerate(inp):
            for x, square in enumerate(row):
                if cls.space(square):
                    inp[y][x] = str(cls.output_of_neighbor_mines(inp, y, x))
        return map(lambda row: "".join(row), inp)

    @classmethod
    def output_of_neighbor_mines(cls, inp, y, x):
        num_mines = cls.num_of_neighbor_mines(inp, y, x)
        return " " if num_mines == 0 else num_mines

    @classmethod
    def num_of_neighbor_mines(cls, inp, y, x):
        return cls.neighbors(inp, y, x).count(cls.MINE)

    @classmethod
    def neighbors(cls, inp, y, x):
        return map(lambda neighbor: inp[neighbor[0]][neighbor[1]],
                   cls.neighbor_coords(inp, y, x))

    @classmethod
    def neighbor_coords(cls, inp, y, x):
        return filter(lambda neighbor: cls.valid_neighbor(inp, neighbor),
                      cls.all_neighbor_coords(inp, y, x))

    @classmethod
    def valid_neighbor(cls, inp, neighbor):
        y, x = neighbor[0], neighbor[1]
        return (0 < y < len(inp) and 0 < x < len(inp[0]) and
                cls.valid_non_border(inp[y][x]))

    @classmethod
    def all_neighbor_coords(cls, inp, y, x):
        return [(y + dy, x + dx) for dy in range(-1, 2) for dx in range(-1, 2)
                if dy != 0 or dx != 0]

    @classmethod
    def valid(cls, inp):
        return (cls.valid_len(inp) and
                cls.valid_border(inp) and
                cls.valid_squares(inp))

    @classmethod
    def valid_len(cls, inp):
        return all(len(row) == len(inp[0]) for row in inp)

    @classmethod
    def valid_border(cls, inp):
        return (cls.valid_middle_borders(inp) and
                cls.valid_first_and_last_borders(inp))

    @classmethod
    def valid_middle_borders(cls, inp):
        return all(cls.valid_middle_border(row) for row in inp[1:-1])

    @classmethod
    def valid_middle_border(cls, row):
        return (cls.vertical_edge(row[0]) and
                cls.vertical_edge(row[-1]))

    @classmethod
    def valid_first_and_last_borders(cls, inp):
        return (cls.valid_first_or_last_border(inp[0]) and
                cls.valid_first_or_last_border(inp[-1]))

    @classmethod
    def valid_first_or_last_border(cls, row):
        return (cls.corner(row[0]) and cls.corner(row[-1]) and
                all(cls.horizontal_edge(square) for square in row[1:-1]))

    @classmethod
    def valid_squares(cls, inp):
        return all(cls.valid_square(square) for row in inp for square in row)

    @classmethod
    def valid_square(cls, square):
        return (cls.mine(square) or
                cls.space(square) or
                cls.corner(square) or
                cls.vertical_edge(square) or
                cls.horizontal_edge(square))

    @classmethod
    def valid_non_border(cls, square):
        return cls.mine(square) or cls.space(square)

    @classmethod
    def mine(cls, square):
        return square == cls.MINE

    @classmethod
    def space(cls, square):
        return square == cls.SPACE

    @classmethod
    def corner(cls, square):
        return square == cls.CORNER

    @classmethod
    def vertical_edge(cls, square):
        return square == cls.VERTICAL_EDGE

    @classmethod
    def horizontal_edge(cls, square):
        return square == cls.HORIZONTAL_EDGE


def board(inp):
    return Minesweeper.board(inp)
