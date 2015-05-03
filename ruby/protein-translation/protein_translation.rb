class Translation

  MAPPINGS = {
    'AUG' => 'Methionine',
    'UUU' => 'Phenylalanine',
    'UUC' => 'Phenylalanine',
    'UUA' => 'Leucine',
    'UUG' => 'Leucine',
    'UCU' => 'Serine',
    'UCC' => 'Serine',
    'UCA' => 'Serine',
    'UCG' => 'Serine',
    'UAU' => 'Tyrosine',
    'UAC' => 'Tyrosine',
    'UGU' => 'Cystine',
    'UGC' => 'Cystine',
    'UGG' => 'Tryptophan',
    'UAA' => 'STOP',
    'UAG' => 'STOP',
    'UGA' => 'STOP'
  }

  def self.of_codon(codon)
    raise InvalidCodonError unless MAPPINGS.include?(codon)
    MAPPINGS[codon]
  end

  def self.of_rna(strand)
    strand.chars.each_slice(3).with_object([]) do |codon, result|
      return result if of_codon(codon.join) == 'STOP'
      result << of_codon(codon.join)
    end
  end

end

class InvalidCodonError < StandardError; end
