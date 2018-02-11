# Nucleotide
class Nucleotide
  NUCLEOTIDES = %w[A T C G].freeze
  attr_reader :strand

  def self.from_dna(strand)
    Nucleotide.new(strand)
  end

  def initialize(strand)
    raise ArgumentError if strand =~ /[^ATCG]/
    @strand = strand
  end

  def histogram
    strand.chars.each_with_object(empty_histogram) do |nucleotide, counts|
      counts[nucleotide] += 1
    end
  end

  def count(nucleotide)
    histogram[nucleotide]
  end

  private

  def empty_histogram
    Hash[NUCLEOTIDES.map { |nucleotide| [nucleotide, 0] }]
  end
end
