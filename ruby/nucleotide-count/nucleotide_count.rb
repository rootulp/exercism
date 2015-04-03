class Nucleotide

  def self.from_dna(strand)
    Nucleotide.new(strand)
  end

  attr_reader :histogram
  def initialize(strand)
    @strand = strand
    @histogram = { 'A' => 0, 'T' => 0, 'C' => 0, 'G' => 0 }
    build_histogram
  end

  def count(letter)
    @histogram[letter]
  end

  private

  def build_histogram
    @strand.each_char do |char|
      raise ArgumentError unless is_nucleotide?(char)
      @histogram[char] += 1
    end
  end

  def is_nucleotide?(char)
    char == 'A' || char == 'T' || char == 'C' || char == 'G' ? true : false
  end
end
