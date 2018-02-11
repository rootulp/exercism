# Matrix
class Matrix
  attr_reader :rows, :columns
  def initialize(string)
    @rows = build_rows(string)
    @columns = rows.transpose
  end

  private

  def build_rows(string)
    string.split("\n").map do |row|
      row.split(' ').map(&:to_i)
    end
  end
end
