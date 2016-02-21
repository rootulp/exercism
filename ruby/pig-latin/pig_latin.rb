require 'set'

# Pig Latin
class PigLatin
  class << self
    ALPHA = ('a'..'z').to_set.freeze
    VOWELS = %w(a e i o u).to_set.freeze
    CONSONANTS = (ALPHA - VOWELS).freeze

    def translate(phrase)
      phrase.split.map { |word| translate_word(word) }.join(' ')
    end

    def translate_word(word)
      return translate_edge_case(word) if edge_case?(word)
      return translate_vowel(word) if starts_with_vowel?(word)
      return translate_consonant(word) if starts_with_consonant?(word)
    end

    def translate_edge_case(word)
      word[3..-1] + word[0..2] + 'ay'
    end

    def translate_vowel(word)
      word + 'ay'
    end

    def translate_consonant(word)
      if starts_with_consonant_sound?(word)
        word[2..-1] + word[0..1] + 'ay'
      else
        word[1..-1] + word[0] + 'ay'
      end
    end

    def edge_case?(word)
      word.start_with?('squ', 'sch', 'thr')
    end

    def starts_with_vowel?(word)
      word.start_with?(*VOWELS, 'yt', 'xr')
    end

    def starts_with_consonant?(word)
      word.start_with?(*CONSONANTS)
    end

    def starts_with_consonant_sound?(word)
      (CONSONANTS.include?(word[0]) && CONSONANTS.include?(word[1])) ||
        word.start_with?('qu')
    end
  end

  private_class_method :translate_edge_case,
                       :translate_vowel,
                       :translate_consonant,
                       :edge_case?,
                       :starts_with_vowel?,
                       :starts_with_consonant?,
                       :starts_with_consonant_sound?
end
