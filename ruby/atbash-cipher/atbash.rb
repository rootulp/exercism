# Atbash Cipher
class Atbash
  ENCODING = Hash[('a'..'z').zip(('a'..'z').to_a.reverse)]

  def self.encode(str)
    split_every_five(encode_word(strip_non_alphanumeric(str.downcase)))
  end

  def self.encode_word(word)
    word.chars.map { |char| encode_char(char) }.join
  end

  def self.encode_char(char)
    alpha?(char) ? ENCODING[char] : char
  end

  def self.strip_non_alphanumeric(str)
    str.gsub(/[^a-z0-9]/i, '')
  end

  def self.split_every_five(str)
    str.scan(/.{5}|.+/).join(' ')
  end

  def self.alpha?(char)
    char.match(/[[:alpha:]]/)
  end

  private_class_method :encode_word, :encode_char, :strip_non_alphanumeric,
                       :split_every_five, :alpha?
end
