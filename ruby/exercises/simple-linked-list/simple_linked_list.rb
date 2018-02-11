# Element
class Element
  attr_reader :datum, :next
  def initialize(datum, next_datum)
    @datum = datum
    @next = next_datum
  end

  def reverse
    from_a(to_a.reverse)
  end

  def to_a
    self.class.to_a(self)
  end

  def from_a(input)
    self.class.from_a(input)
  end

  class << self
    def to_a(curr)
      result = []
      while curr.is_a?(Element)
        result << curr.datum
        curr = curr.next
      end
      result
    end

    def from_a(input)
      arr = input.to_a
      return nil if arr.empty? || !arr.is_a?(Array)

      result = Element.new(arr.pop, nil)
      result = Element.new(arr.pop, result) while arr.any?
      result
    end
  end
end
