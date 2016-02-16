# Triangle
class Triangle
  attr_reader :num_rows
  def initialize(num_rows)
    @num_rows = num_rows
  end

  def rows
    (0..(num_rows - 1)).map { |row_num| build_row(row_num) }
  end

  def build_row(row_num)
    (0..row_num).map { |col_num| n_choose_k(row_num, col_num) }
  end

  private

  def n_choose_k(n, k)
    factorial(n) / (factorial(k) * factorial(n - k))
  end

  def factorial(n)
    return 1 if n == 0
    n.downto(1).reduce(:*)
  end
end
