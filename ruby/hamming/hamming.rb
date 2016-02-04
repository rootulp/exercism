# Hamming
class Hamming
  def self.compute(s1, s2)
    return compute_char(s1[0], s2[0]) if s1.size == 1 || s2.size == 1
    compute_char(s1[0], s2[0]) + compute(s1[1..-1], s2[1..-1])
  end

  def self.compute_char(c1, c2)
    c1 == c2 ? 0 : 1
  end
  private_class_method :compute_char
end
