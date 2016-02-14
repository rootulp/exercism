# Nucleotide
class Nucleotide
  def self.from_dna(strand)
    Nucleotide.new(strand)
  end

  attr_reader :strand, :histogram
  def initialize(strand)
    @strand = strand
    @histogram = { 'A' => 0, 'T' => 0, 'C' => 0, 'G' => 0 }
    build_histogram
  end

  def count(nucleotide)
    histogram[nucleotide]
  end

  private

  def build_histogram
    strand.each_char do |symbol|
      fail ArgumentError unless nucleotide?(symbol)
      histogram[symbol] += 1
    end
  end

  def nucleotide?(symbol)
    symbol == 'A' || symbol == 'T' || symbol == 'C' || symbol == 'G'
  end
end
