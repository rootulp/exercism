# Phrase
class Phrase
  attr_reader :phrase
  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    words.each_with_object(Hash.new(0)) { |word, counts| counts[word] += 1 }
  end

  private

  def words
    phrase.downcase.gsub(/[^'a-z0-9\s]/i, ' ').split(' ')
  end
end
