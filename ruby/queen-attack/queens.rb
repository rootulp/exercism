class Queens
  attr_accessor :board
  attr_reader :white, :black

  DEFAULTS = {white: [0, 3], black: [7, 3]}

  def initialize(args=DEFAULTS)
    raise ArgumentError if args[:white] == args[:black]

    @board = Array.new(8) { Array.new(8) }
    @white = [args[:white][0], args[:white][1]]
    @black = [args[:black][0], args[:black][1]]
    set_queens
  end

  def attack?
    same_row? || same_col? || same_diag?
  end

  private

  def set_queens
    board[white[0], white[1]] = "W"
    board[black[0], black[1]] = "B"
  end

  def same_row?
    white[0] == black[0]
  end

  def same_col?
    white[1] == black[1]
  end

  def same_diag?
    (white[1] - white[0]).abs == (black[1] - black[0]).abs
  end

end
