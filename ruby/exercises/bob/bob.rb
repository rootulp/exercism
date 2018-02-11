# Bob
class Bob
  def hey(phrase)
    if phrase.strip.empty?
      'Fine. Be that way!'
    elsif phrase == phrase.upcase && phrase != phrase.downcase
      'Whoa, chill out!'
    elsif phrase.end_with?('?')
      'Sure.'
    else
      'Whatever.'
    end
  end
end
