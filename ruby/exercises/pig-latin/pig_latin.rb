require 'set'

# Pig Latin
class PigLatin
  class << self
    ALPHA = ('a'..'z').to_set.freeze
    VOWELS = %w[a e i o u].to_set.freeze
    CONSONANTS = (ALPHA - VOWELS).freeze

    def translate(phrase)
      phrase.split.map { |word| translate_word(word) }.join(' ')
    end

    def translate_word(word)
      return vowel(word) if vowel_case?(word)
      return edge_case(word) if edge_case?(word)
      return consonant_sound(word) if consonant_sound_case?(word)
      consonant(word)
    end

    def vowel(word)
      word + 'ay'
    end

    def edge_case(word)
      word[3..-1] + word[0..2] + 'ay'
    end

    def consonant_sound(word)
      word[2..-1] + word[0..1] + 'ay'
    end

    def consonant(word)
      word[1..-1] + word[0] + 'ay'
    end

    def vowel_case?(word)
      word.start_with?(*VOWELS, 'yt', 'xr')
    end

    def edge_case?(word)
      word.start_with?('squ', 'sch', 'thr')
    end

    def consonant_sound_case?(word)
      (CONSONANTS.include?(word[0]) && CONSONANTS.include?(word[1])) ||
        word.start_with?('qu')
    end
  end

  private_class_method :vowel,
                       :edge_case,
                       :consonant_sound,
                       :consonant,
                       :vowel_case?,
                       :edge_case?,
                       :consonant_sound_case?
end
