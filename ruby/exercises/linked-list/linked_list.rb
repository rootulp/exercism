# Deque
class Deque
  # Element
  Element = Struct.new(:data, :next, :prev) do
  end

  attr_reader :head, :tail
  def initialize
    reset!
  end

  def push(data)
    new_element = Element.new(data)
    return @head = new_element && @tail = new_element if empty?
    new_element.prev = tail
    tail.next = new_element
    @tail = new_element
  end

  def unshift(data)
    new_element = Element.new(data)
    return @head = new_element && @tail = new_element if empty?
    new_element.next = head
    head.prev = new_element
    @head = new_element
  end

  def pop
    prev_tail = tail
    if one_element?
      reset!
    else
      @tail = prev_tail.prev
      prev_tail.next = false
    end
    prev_tail.data
  end

  def shift
    prev_head = head
    if one_element?
      reset!
    else
      @head = prev_head.next
      prev_head.prev = false
    end
    prev_head.data
  end

  private

  def reset!
    @head = false
    @tail = false
  end

  def one_element?
    head == tail
  end

  def empty?
    !head && !tail
  end
end
