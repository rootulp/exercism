# Solution from https://github.com/Azdaroth/exercism.io/blob/master/ruby/accumulate/array.rb
class Array
  def accumulate
    accumulated = []

    each do |element|
      accumulated << (yield element)
    end

    accumulated
  end
end
