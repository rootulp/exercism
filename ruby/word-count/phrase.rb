class Phrase
  attr_reader :word_count

  def initialize(sentence)
    words = extract_words(sentence)
    @word_count = count(words)
  end

  def extract_words(sentence)
    sentence.downcase.gsub(/[^'a-z0-9\s]/i, " ").split(" ")
  end

  def count(words)
    occurences = Hash.new 0
    words.each { |word| occurences[word] += 1 }
    occurences
  end

end