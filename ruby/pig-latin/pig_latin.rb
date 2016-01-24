class PigLatin
  require 'set'

  ALPHA = ('a'..'z').to_set
  VOWELS = ['a', 'e', 'i', 'o', 'u'].to_set
  CONSONANTS = ALPHA - VOWELS

  def self.translate(phrase)
    phrase.split.map { |word| self.translate_word(word) }.join(' ')
  end

  def self.translate_word(word)
    if VOWELS.include?(word[0]) ||
         word.start_with?('yt') ||
         word.start_with?('xr')
      word + 'ay'

    elsif word.start_with?('squ') ||
          word.start_with?('sch') ||
          word.start_with?('thr')
      word[3..-1] + word[0..2] + 'ay'

    elsif (CONSONANTS.include?(word[0])  &&
           CONSONANTS.include?(word[1])) ||
           word.start_with?('qu')
      word[2..-1] + word[0..1] + 'ay'

    elsif CONSONANTS.include?(word[0])
      word[1..-1] + word[0..0] + 'ay'

    end
  end
end
