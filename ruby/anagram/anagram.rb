class Anagram

  attr_reader :init_word, :chars

  def initialize(init_word)
    @init_word = init_word
    @chars = init_word.downcase.split(//).sort!
  end

  def match(words)
    anagrams = []
    words.each do |word|
      anagrams << word if is_anagram?(word)
    end
    anagrams
  end

  def is_anagram?(word)
    return false if word.downcase == init_word
    word.downcase.split(//).sort! == chars ? true : false
  end

end