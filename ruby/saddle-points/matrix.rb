class Matrix

  attr_reader :rows, :columns, :saddle_points
  def initialize(string)
    @rows = build_rows(string)
    @columns = rows.transpose
    @saddle_points = find_saddle_points
  end

  private

  def build_rows(string)
    string.split("\n").map do |row|
      row.split(" ").map(&:to_i)
    end
  end

  def find_saddle_points
    results = []
    rows.each_with_index do |row, x|
      row.each_with_index do |val, y|
        results << [x,y] if saddle_point?(val, x, y)
      end
    end
    results
  end

  def saddle_point?(val, x, y)
    max_x = rows[x].max
    min_y = columns[y].min
    val == max_x && val == min_y 
  end
end
