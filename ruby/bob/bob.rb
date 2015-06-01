class Bob

  def hey(phrase)
    if phrase.strip.empty?
      return "Fine. Be that way!"
    elsif phrase == phrase.upcase and phrase != phrase.downcase
      return "Whoa, chill out!"
    elsif phrase.end_with?("?")
      return "Sure."
    else
      return 'Whatever.'
    end
  end

end
