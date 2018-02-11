# Anagram analyzes a list of words for anagrams of a given word
class Anagram
  attr_reader :original_word, :original_sorted_chars
  def initialize(word)
    @original_word = word.downcase
    @original_sorted_chars = original_word.chars.sort
  end

  def match(words)
    words.select { |word| anagram?(word.downcase) }
  end

  private

  def anagram?(word)
    word != original_word && word.chars.sort == original_sorted_chars
  end
end
