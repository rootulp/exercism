class Deque

  def initialize
    @head = nil
    @tail = nil
  end

  def push(val)
    new = Node.new(val)

    if @head.nil? && @tail.nil?
      @head = new
      @tail = new
    elsif @tail.nil?
      @tail = new
    else
      new.prev = @tail
      @tail.next = new
      @tail = new
    end
  end

  def pop
    if @tail == @head
      res = @tail
      @tail = nil
      @head = nil
      return res.val
    else
      res = @tail
      @tail = res.prev
      @tail.next = nil
      return res.val
    end
  end

  def unshift(val)
    new = Node.new(val)

    if @head.nil? && @tail.nil?
      @head = new
      @tail = new
    elsif @head.nil?
      @head = new
    else
      new.next = @head
      @head.prev = new
      @head = new
    end
  end

  def shift
    if @tail == @head
      res = @head
      @tail = nil
      @head = nil
      return res.val
    else
      res = @head
      @head = res.next
      @head.prev = nil
      return res.val
    end
  end
end

class Node
  attr_reader :val
  attr_accessor :next, :prev
  def initialize(val)
    @val = val
    @next = nil
    @prev = nil
  end
end
