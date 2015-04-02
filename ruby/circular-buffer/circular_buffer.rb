class CircularBuffer

  class BufferEmptyException < StandardError
  end

  class BufferFullException < StandardError
  end

  attr_accessor :buffer, :size, :head, :tail
  def initialize(size)
    @size = size
    clear
  end

  def clear
    @buffer = Array.new(size, nil)
    @head = 0
    @tail = 0
  end

  def write!(val)
    write_buffer(val)
    increment_tail
  end

  def write(val)
    raise BufferFullException if head == tail && buffer[head] != nil
    write_buffer(val)
  end

  def read
    raise BufferEmptyException if buffer[tail] == nil
    result = buffer[tail]
    buffer[tail] = nil
    increment_tail
    return result
  end

  private

  def write_buffer(val)
    return if val == nil
    buffer[head] = val
    self.head = (head + 1) % size
  end

  def increment_tail
    self.tail = (tail + 1) % size
  end

end
