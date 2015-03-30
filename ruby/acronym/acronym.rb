class Acronym

  def self.abbreviate(input)
    Acronym.new(input).abbreviate
  end

  attr_reader :input
  def initialize(input)
    @input = input
  end

  def abbreviate
    return extract_acronym if is_acronym?
    words = input.split(/(?=[A-Z])|\W+/)
    words.map { |word| word[0] }.join.upcase
  end

  private

  def is_acronym?
    input.include?(':')
  end

  def extract_acronym
    input.partition(':').first
  end

end
