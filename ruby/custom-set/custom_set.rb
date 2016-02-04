# Custom Set
class CustomSet
  attr_reader :arr
  def initialize(vals = [])
    @arr = []
    put_all(vals.to_a)
  end

  def put(val)
    new_node = Node.new(val)
    arr << new_node unless arr.include?(new_node)
    arr.sort!
    self
  end

  def delete(val)
    other_node = Node.new(val)
    arr.delete_if { |node| node == other_node }
    self
  end

  def difference(other_set)
    other_set.arr.each { |other_node| delete(other_node.val) }
    self
  end

  def union(other_set)
    other_set.arr.each { |other_node| put(other_node.val) }
    self
  end

  def intersection(other_set)
    shared = []
    other_set.arr.each do |other_node|
      shared << other_node.val if arr.include?(other_node)
    end
    CustomSet.new(shared)
  end

  def disjoint?(other_set)
    arr.each { |node| return false if other_set.arr.include?(node) }
    true
  end

  def subset?(other_set)
    other_set.arr.each do |other_node|
      return false unless arr.include?(other_node)
    end
    true
  end

  def member?(val)
    new_node = Node.new(val)
    arr.member?(new_node)
  end

  def put_all(vals)
    vals.each { |val| put(val) }
  end

  def to_a
    arr.map(&:val)
  end

  def ==(other)
    arr == other.arr
  end

  def empty
    arr.clear
    self
  end

  def size
    arr.size
  end
end

# Node
class Node
  include Comparable

  attr_reader :val
  def initialize(val)
    @val = val
  end

  def ==(other)
    val == other.val && val.class == other.val.class
  end

  def <=>(other)
    val <=> other.val
  end
end
