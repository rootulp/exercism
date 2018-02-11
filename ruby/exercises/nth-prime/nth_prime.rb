# NthPrime
class NthPrime
  class << self
    def nth(num)
      fail ArgumentError if num < 1
      primes(num).last
    end

    def primes(num)
      found = []
      curr = 1
      while found.size < num
        found << curr if prime?(curr)
        curr += 1
      end
      found
    end

    def prime?(num)
      return false if num == 1
      return true if num == 2
      (2..Math.sqrt(num).ceil).each do |i|
        return false if num % i == 0
      end
      true
    end
  end
end
