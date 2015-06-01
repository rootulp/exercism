class Hamming

  def self.compute(strand1, strand2)
    count = 0

    strand1 = strand1.split(//)
    strand2 = strand2.split(//)

    strand1.each_with_index do |char,index|
      count += 1 if strand2[index] != char
    end

    return count

  end

end
