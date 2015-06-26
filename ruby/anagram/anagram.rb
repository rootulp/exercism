class Anagram

  attr_reader :detector_word, :detector_chars
  def initialize(w)
    @detector_word = w.downcase
    @detector_chars = detector_word.chars.sort
  end

  def match(words)
    words.select { |word| is_anagram?(word.downcase) }
  end

  private

  def is_anagram?(word)
    word != detector_word ? word.chars.sort == detector_chars : false
  end

end
