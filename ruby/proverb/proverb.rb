class Proverb

  attr_reader :items, :qualifier
  def initialize(*items, qualifier: nil)
    @items = items
    @qualifier = "#{qualifier} #{items.first}".strip
  end

  def to_s
    output = String.new
    items.each_cons(2) do |item1, item2|
      output << partial(item1, item2)
    end
    output << ending
  end

  private

  def partial(item1, item2)
    "For want of a #{item1} the #{item2} was lost.\n" 
  end

  def ending
    "And all for the want of a #{qualifier}."
  end

end
