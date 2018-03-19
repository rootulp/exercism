# Saddle Points
class SaddlePoints
  attr_reader :input
  def initialize(input)
    @input = input
  end

  def rows
    input.split("\n").map do |row|
      row.split(' ').map(&:to_i)
    end
  end

  def columns
    rows.transpose
  end

  def saddle_points
    results = []
    rows.each_with_index do |row, x|
      row.each_with_index do |val, y|
        results << [x, y] if saddle_point?(val, x, y)
      end
    end
    results
  end

  # rubocop:disable Naming/UncommunicativeMethodParamName
  def saddle_point?(val, x, y)
    val == rows[x].max && val == columns[y].min
  end
  # rubocop:enable Naming/UncommunicativeMethodParamName
end
