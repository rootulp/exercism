require 'prime'

class Prime

  def nth(num)
    raise ArgumentError if num < 1
    Prime.take(num)[num-1]
  end

end