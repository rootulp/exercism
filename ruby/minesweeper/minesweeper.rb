# Board
class Board
  def self.transform(input)
    Board.new(input).transform
  end

  attr_reader :grid
  def initialize(input)
    @grid = input.map { |row| row.split(//) }
    fail ValueError if error?
  end

  def transform
    grid.each_with_index.map do |r, row|
      r.each_with_index.map do |square, col|
        occupied?(square) ? square : mines_for(row, col)
      end.join
    end
  end

  private

  def neighbors(row, col)
    neighbor_positions.map { |dx, dy| grid[row + dx][col + dy] }
  end

  def neighbor_positions
    (-1..1).to_a.product((-1..1).to_a).reject { |dx, dy| dx == 0 && dy == 0 }
  end

  def mines_for(row, col)
    return ' ' if neighbors(row, col).count { |square| mine?(square) } == 0
    neighbors(row, col).count { |square| mine?(square) }.to_s
  end

  def occupied?(square)
    border?(square) || mine?(square)
  end

  def border?(square)
    square == '|' || square == '-' || square == '+'
  end

  def mine?(square)
    square == '*'
  end

  def error?
    invalid_row_length? || invalid_border? || invalid_square?
  end

  def invalid_border?
    invalid_border_row?(grid.first) ||
      invalid_border_row?(grid.last) ||
      invalid_border_edges?
  end

  def invalid_border_row?(row)
    row.any? { |square| !border?(square) }
  end

  def invalid_border_edges?
    grid.any? { |row| !border?(row.first) || !border?(row.last) }
  end

  def invalid_row_length?
    grid.any? { |row| row.length != grid.first.length }
  end

  def invalid_square?
    grid.join =~ /[^\s*|+-]/
  end
end

class ValueError < StandardError
end
