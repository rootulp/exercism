class CustomSet

  attr_accessor :arr
  def initialize(vals = [])
    @arr = populate_arr(vals.to_a)
  end

  def populate_arr(vals)
    arr = []
    vals.uniq.sort.each do |val|
      arr << Node.new(val)
    end
    arr
  end

  def ==(other_set)
    arr.zip(other_set.arr).each do |node1, node2|
      return false if node1.val != node2.val
    end
    true
  end

  def delete(val)
    other_node = Node.new(val)
    arr.delete_if do |node|
      node == other_node
    end
    self
  end

  def difference(other_set)
    other_set.arr.each do |other_node|
      delete(other_node.val)
    end
    self
  end

  def disjoint?(other_set)
    arr.each do |node|
      return false if other_set.arr.include?(node)
    end
    true
  end

  def empty
    arr.clear
    self
  end

  def intersection(other_set)
    shared = []
    other_set.arr.each do |other_node|
      shared << other_node.val if arr.include?(other_node)
    end
    CustomSet.new(shared)
  end

  def member?(val)
    new_node = Node.new(val)
    arr.each do |node|
      return true if node == new_node
    end
    false
  end

  def put(val)
    new_node = Node.new(val)
    arr << new_node unless arr.include?(new_node)
    arr.sort!
    self
  end

  def size
    arr.size
  end

  def subset?(other_set)
    other_set.arr.each do |other_node|
      return false unless arr.include?(other_node)
    end
    true
  end

  def to_a
    arr.map { |node| node.val }
  end

  def union(other_set)
    other_set.arr.each do |node|
      put(node.val)
    end
    self
  end

end

class Node
  include Comparable

  attr_reader :val
  def initialize(val)
    @val = val
  end

  def ==(other_node)
    val == other_node.val && val.class == other_node.val.class
  end

  def <=>(other_node)
    val <=> other_node.val
  end
end
