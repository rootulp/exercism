class Atbash
  ENCODING = Hash[('a'..'z').zip('a'..'z').to_a.reverse]

  def self.encode(str)
    split_5(encode!(format(str)))
  end

  def self.encode!(str)
    str.chars.map { |char| is_numeric?(char) ? char : ENCODING[char] }.join('')
  end

  def self.format(str)
    str.downcase.gsub(/[^a-z0-9]/i, '')
  end

  def self.split_5(str)
    str.scan(/.{5}|.+/).join(' ')
  end

  # Taken from Rosetta code
  def self.is_numeric?(s)
    !!Float(s) rescue false
  end
end
