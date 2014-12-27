class Complement

  def self.of_dna(strand)
    result = ""
    strand.each_char do |nucleobase|
      result << self.dna_pairing_for(nucleobase)
    end

    result
  end

  def self.of_rna(strand)
    result = ""
    strand.each_char do |nucleobase|
      result << self.rna_pairing_for(nucleobase)
    end

    result
  end

  def self.dna_pairing_for(nucleobase)
    pairings = {
      "G" => "C",
      "C" => "G",
      "A" => "U",
      "T" => "A"
    }

    pairings[nucleobase]
  end

  def self.rna_pairing_for(nucleobase)
    pairings = {
      "G" => "C",
      "C" => "G",
      "A" => "T",
      "U" => "A"
    }

    pairings[nucleobase]
  end

end