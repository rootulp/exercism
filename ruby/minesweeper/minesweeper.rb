class Board

  def self.transform(inp)
    Board.new(inp).transform
  end

  attr_reader :grid
  def initialize(inp)
    @grid = inp.map { |row| row.split(//) }
    raise ValueError if error?
  end

  def transform
    grid.each_with_index.map do |r, row|
      r.each_with_index.map do |val, col|
        reserved?(val) ? val : mines_for(row, col)
      end.join
    end
  end

  private

  def neighbors(row, col)
    neighbors = []
    (-1..1).each do |x|
      (-1..1).each do |y|
        next if x == 0 && y == 0
        neighbors << grid[row + x][col + y]
      end
    end
    neighbors
  end

  def mines_for(row, col)
    count = neighbors(row, col).count('*')
    count == 0 ? " " : "#{count}"
  end

  def reserved?(val)
    border?(val) || val == '*'
  end

  def border?(val)
    val == '|' || val == '-' || val == '+'
  end

  # Error Checking

  def error?
    different_len? || faulty_border? || invalid_char?
  end

  def faulty_border?
    invalid_border_row?(grid.first) ||
    invalid_border_row?(grid.last)  ||
    invalid_border_edges?
  end

  def invalid_border_row?(row)
    row.each do |val|
      return true unless border?(val)
    end
    false
  end

  def invalid_border_edges?
    grid.each do |row|
      return true unless border?(row.first) && border?(row.last)
    end
    false
  end

  def different_len?
    size = grid.first.length
    grid.each do |row|
      return true if row.length != size
    end
    false
  end

  def invalid_char?
    grid.join =~ /[^\s*|+-]/
  end

end

class ValueError < StandardError
end
