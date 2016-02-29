class Complement
  DNA_PAIRINGS = {
    'G' => 'C',
    'C' => 'G',
    'A' => 'U',
    'T' => 'A'
  }.freeze
  RNA_PAIRINGS = {
    'G' => 'C',
    'C' => 'G',
    'A' => 'T',
    'U' => 'A'
  }.freeze
  def self.of_dna(strand)
    raise ArgumentError if strand.include?('U')
    strand.split(//).map { |nucleobase| dna_pairing_for(nucleobase) }.join
  end

  def self.of_rna(strand)
    raise ArgumentError if strand.include?('T')
    strand.split(//).map { |nucleobase| rna_pairing_for(nucleobase) }.join
  end

  def self.dna_pairing_for(nucleobase)
    DNA_PAIRINGS[nucleobase]
  end

  def self.rna_pairing_for(nucleobase)
    RNA_PAIRINGS[nucleobase]
  end
end
