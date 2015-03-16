# Adapted from
# https://github.com/alxndr/exercism/blob/master/ruby/house/house.rb

class House

  VERSE_MAPPING = {
    12=> {subject: 'horse and the hound and the horn', verb: 'belonged to' },
    11=> {subject: 'farmer sowing his corn',           verb: 'kept' },
    10=> {subject: 'rooster that crowed in the morn',  verb: 'woke' },
    9 => {subject: 'priest all shaven and shorn',      verb: 'married' },
    8 => {subject: 'man all tattered and torn',        verb: 'kissed' },
    7 => {subject: 'maiden all forlorn',               verb: 'milked' },
    6 => {subject: 'cow with the crumpled horn',       verb: 'tossed' },
    5 => {subject: 'dog',                              verb: 'worried' },
    4 => {subject: 'cat',                              verb: 'killed' },
    3 => {subject: 'rat',                              verb: 'ate' },
    2 => {subject: 'malt',                             verb: 'lay in the house'\
                                                             ' that Jack built'}
  }

  def self.verse(num)
    return "This is the house that Jack built.\n" if num == 1

    lines = num.downto(2).map do |n|
      verse_parts = VERSE_MAPPING[n]
      "#{verse_parts[:subject]}\nthat #{verse_parts[:verb]}"
    end

    "This is the #{lines.join(' the ')}.\n"
  end

  def self.verses(from, to)
    ( from.upto(to).map { |n| self.verse(n) }.join("\n"))
  end

  def self.recite
    verses(1,12)
  end

end
