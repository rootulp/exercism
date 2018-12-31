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
        return cls.solve(inp)

    # Split rows (String -> List) and operate on board in place
    @classmethod
    def solve(cls, inp):
        inp = list(map(lambda row: list(row), inp))
        return list(map(lambda row: "".join(row), cls.generate_board(inp)))

    @classmethod
    def generate_board(cls, inp):
        return [[cls.convert_square(inp, y, x)
                 for x, square in enumerate(row)]
                for y, row in enumerate(inp)]

    # Only convert squares that are spaces
    @classmethod
    def convert_square(cls, inp, y, x):
        if not cls.is_space(inp[y][x]):
            return inp[y][x]
        return str(cls.output_of_neighbor_mines(inp, y, x))

    @classmethod
    def output_of_neighbor_mines(cls, inp, y, x):
        num_mines = cls.num_of_neighbor_mines(inp, y, x)
        return " " if num_mines == 0 else num_mines

    @classmethod
    def num_of_neighbor_mines(cls, inp, y, x):
        return len(
            list(filter(
                lambda neighbor: cls.is_neighbor_a_mine(
                    inp, neighbor), cls.all_neighbor_coords(
                    inp, y, x))))

    # Checks if coords are within bounds then checks for is_mine
    @classmethod
    def is_neighbor_a_mine(cls, inp, neighbor):
        y, x = neighbor[0], neighbor[1]
        return (0 < y < len(inp) and 0 < x < len(inp[0]) and
                cls.is_mine(inp[y][x]))

    # Generates list of tuples for all neighboring coords
    # (excluding current coord)
    @classmethod
    def all_neighbor_coords(cls, inp, y, x):
        return [(y + dy, x + dx) for dy in range(-1, 2) for dx in range(-1, 2)
                if dy != 0 or dx != 0]

    @classmethod
    def valid(cls, inp):
        return (cls.valid_len(inp) and
                cls.valid_border(inp) and
                cls.valid_squares(inp))

    # Tests if all rows are the same size
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
        return (cls.is_vertical_edge(row[0]) and
                cls.is_vertical_edge(row[-1]))

    @classmethod
    def valid_first_and_last_borders(cls, inp):
        return (cls.valid_first_or_last_border(inp[0]) and
                cls.valid_first_or_last_border(inp[-1]))

    @classmethod
    def valid_first_or_last_border(cls, row):
        return (cls.is_corner(row[0]) and cls.is_corner(row[-1]) and
                all(cls.is_horizontal_edge(square) for square in row[1:-1]))

    @classmethod
    def valid_squares(cls, inp):
        return all(cls.valid_square(square) for row in inp for square in row)

    @classmethod
    def valid_square(cls, square):
        return (cls.is_mine(square) or
                cls.is_space(square) or
                cls.is_corner(square) or
                cls.is_vertical_edge(square) or
                cls.is_horizontal_edge(square))

    @classmethod
    def valid_non_border(cls, square):
        return cls.is_mine(square) or cls.is_space(square)

    @classmethod
    def is_mine(cls, square):
        return square == cls.MINE

    @classmethod
    def is_space(cls, square):
        return square == cls.SPACE

    @classmethod
    def is_corner(cls, square):
        return square == cls.CORNER

    @classmethod
    def is_vertical_edge(cls, square):
        return square == cls.VERTICAL_EDGE

    @classmethod
    def is_horizontal_edge(cls, square):
        return square == cls.HORIZONTAL_EDGE


def board(inp):
    return Minesweeper.board(inp)
