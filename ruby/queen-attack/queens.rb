# Queens
class Queens
  DEFAULTS = { white: [0, 3], black: [7, 3] }.freeze

  attr_reader :white, :black
  attr_accessor :board
  def initialize(positions = DEFAULTS)
    fail ArgumentError if positions[:white] == positions[:black]
    @board = Array.new(8) { Array.new(8, '_') }
    @white = [*positions[:white]]
    @black = [*positions[:black]]
    set_queen('W', white.first, white.last)
    set_queen('B', black.first, black.last)
  end

  def attack?
    same_row? || same_col? || same_diagonal?
  end

  def to_s
    board.map { |row| row.join(' ') }.join("\n")
  end


  private

  def set_queen(color, x, y)
    board[x][y] = color
  end

  # BUG hard-coded indices
  def same_row?
    white[0] == black[0]
  end

  def same_col?
    white[1] == black[1]
  end

  def same_diagonal?
    (white[1] - white[0]).abs == (black[1] - black[0]).abs
  end
end
