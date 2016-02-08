# Square Code Crypto
class Crypto
  attr_reader :unnormalized
  def initialize(unnormalized)
    @unnormalized = unnormalized
  end

  def normalize_plaintext
    unnormalized.gsub(/\W/, '').downcase
  end

  def size
    (normalize_plaintext.size**0.5).ceil
  end

  def plaintext_segments
    normalize_plaintext.scan(/.{1,#{size}}/)
  end

  def ciphertext
    safe_transpose(matrix).join('')
  end

  def normalize_ciphertext
    ciphertext.scan(/.{1,#{matrix.size}}/).join(' ')
  end

  private

  def matrix
    plaintext_segments.map { |segment| segment.split(//) }
  end

  def safe_transpose(matrix)
    matrix.first.zip(*matrix.drop(1))
  end
end
