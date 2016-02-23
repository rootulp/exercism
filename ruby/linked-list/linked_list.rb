# Deque
class Deque
  attr_reader :head, :tail
  def initialize
    reset!
  end

  def push(data)
    new_element = Element.new(data)
    return set_head_and_tail(new_element) if empty?
    new_element.prev = tail
    tail.next = new_element
    set_tail(new_element)
  end

  def unshift(data)
    new_element = Element.new(data)
    return set_head_and_tail(new_element) if empty?
    new_element.next = head
    head.prev = new_element
    set_head(new_element)
  end

  def pop
    prev_tail = tail
    if size_is_one?
      reset!
    else
      set_tail(prev_tail.prev)
      prev_tail.next = false
    end
    prev_tail.data
  end

  def shift
    prev_head = head
    if size_is_one?
      reset!
    else
      set_head(prev_head.next)
      prev_head.prev = false
    end
    prev_head.data
  end

  private

  def empty?
    !head && !tail
  end

  def size_is_one?
    head == tail
  end

  def reset!
    set_head_and_tail(false)
  end

  def set_head_and_tail(element)
    set_head(element)
    set_tail(element)
  end

  def set_head(element)
    @head = element
  end

  def set_tail(element)
    @tail = element
  end
end

# Element
class Element
  attr_reader :data
  attr_accessor :next, :prev
  def initialize(data)
    @data = data
  end
end
