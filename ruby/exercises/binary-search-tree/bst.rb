# Binary Search Tree
class Bst
  attr_reader :data, :left, :right

  def initialize(data)
    @data = data
  end

  def insert(val)
    val <= data ? insert_left(val) : insert_right(val)
  end

  def insert_left(data)
    left ? left.insert(data) : @left = Bst.new(data)
  end

  def insert_right(data)
    right ? right.insert(data) : @right = Bst.new(data)
  end

  def each(&block)
    left.each(&block) if left
    yield data
    right.each(&block) if right
  end
end
