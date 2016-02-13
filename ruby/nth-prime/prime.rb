require 'prime'

# Prime
class Prime
  def nth(num)
    fail ArgumentError if num < 1
    Prime.take(num)[num - 1]
  end
end
