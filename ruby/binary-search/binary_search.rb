class BinarySearch
  
  attr_reader :list
  def initialize(list)
    raise ArgumentError if list != list.sort
    @list = list
  end

  def search_for(val)
    search(val, 0, list.size - 1)
  end

  def middle
    (list.size - 1) / 2
  end

  private

  def search(val, left, right)
    raise RuntimeError if left >= right 
    middle = (right + left) / 2

    if val < list[middle]
      search(val, left, middle - 1)
    elsif val > list[middle]
      search(val, middle + 1, right)
    else
      middle
    end
  end

end
