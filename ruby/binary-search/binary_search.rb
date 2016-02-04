# Performs a Binary Search for a particular int in an array
class BinarySearch
  attr_reader :list
  def initialize(list)
    @list = list
    fail ArgumentError if list != list.sort
  end

  def search_for(val)
    search(val, 0, list.size - 1)
  end

  def middle(left = 0, right = list.size - 1)
    (left + right) / 2
  end

  private

  def search(val, left, right)
    fail RuntimeError if left >= right

    mid = middle(left, right)
    return search(val, left, mid - 1) if val < list[mid]
    return search(val, mid + 1, right) if val > list[mid]
    mid
  end
end
