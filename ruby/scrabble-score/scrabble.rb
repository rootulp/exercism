class Scrabble

  LETTER_VALUES = {
    A: 1, B: 3, C: 3, D: 2, E: 1, F: 4, G: 2, H: 4,
    I: 1, J: 8, K: 5, L: 1, M: 3, N: 1, O: 1, P: 3,
    Q:10, R: 1, S: 1, T: 1, U: 1, V: 4, W: 4, X: 8,
    Y: 4, Z:10
  }

  attr_reader :word
  def initialize(word)
    @word = (word || "").gsub(/\s+/, "").upcase
  end

  def score
    points = 0
    word.each_char {|char| points += LETTER_VALUES[char.to_sym]}
    points
  end

  def self.score(word)
    self.new(word).score
  end

end
