require 'set'

# Pangram
class Pangram
  VERSION = 1

  def self.pangram?(str)
    alphabet = ('a'..'z').to_set
    str.downcase.chars.each do |char|
      alphabet.delete(char)
    end
    alphabet.empty?
  end
  singleton_class.send(:alias_method, :is_pangram?, :pangram?)
end
