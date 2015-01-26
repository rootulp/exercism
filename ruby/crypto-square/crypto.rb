class Array
  def safe_transpose
    result = []
    max_size = self.max { |a,b| a.size <=> b.size }.size
    max_size.times do |i|
      result[i] = Array.new(self.first.size)
      self.each_with_index { |r,j| result[i][j] = r[i] }
    end
    result
  end
end

class Crypto

  attr_reader :unnormalized, :normalized
  def initialize(unnormalized)
    @unnormalized = unnormalized
  end

  def normalize_plaintext
    unnormalized.gsub(/\W/, '').downcase
  end

  def size
    (normalize_plaintext.length ** 0.5).ceil
  end

  def plaintext_segments
    normalize_plaintext.scan(/.{1,#{size}}/)
  end

  def ciphertext
    matrix.safe_transpose.join("")
  end

  def matrix
    plaintext_segments.map{|arr| arr.split(//)}
  end

  def normalize_ciphertext
    ciphertext.scan(/.{1,#{matrix.length}}/).join(" ")
  end

end