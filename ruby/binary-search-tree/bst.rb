class Node

  attr_reader :data
  attr_accessor :left, :right
  def initialize(val)
    @data = val
  end

end

class Bst

  attr_reader :root
  def initialize(val)
    @root = Node.new(val)
  end

  def data
    root.data
  end

  def left
    root.left
  end

  def right
    root.right
  end

  def insert(val)
    insert_value(root, val)
  end

  def insert_value(node, val)
    if node.nil?
      node = Node.new(val)
    elsif val > node.data
      node.right = insert_value(node.right, val)
    else
      node.left = insert_value(node.left, val)
    end
    node
  end

end
