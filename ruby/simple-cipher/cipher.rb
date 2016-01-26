class Cipher

  KEY_LENGTH = 100

  attr_reader :key
  def initialize(key = generate_key)
    raise ArgumentError if invalid_key?(key)
    @key = key
  end

  def encode(text)
    text.chars.map.with_index do |char, index|
      (wrap(char.ord + shift_ord(index))).chr
    end.join
  end

  def decode(text)
    text.chars.map.with_index do |char, index|
      (wrap(char.ord - shift_ord(index))).chr
    end.join
  end

  private

  def generate_key
    (0..KEY_LENGTH).map { ('a'..'z').to_a[rand(26)] }.join
  end

  def shift_ord(index)
    key[index].ord - 97
  end

  def wrap(val)
    val -= 26 if val > 122
    val += 26 if val < 97
    val
  end

  def invalid_key?(key)
    key.empty? ||
    key =~ /\d/ ||
    key =~ /[A-Z]/
  end

end
