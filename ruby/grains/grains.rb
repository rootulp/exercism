class Grains

  def initialize
    @board = (0..63).map {|x| 2 ** x }
  end

  def square(num)
    @board[num-1]
  end

  def total
    @board.inject(:+)
  end

end
