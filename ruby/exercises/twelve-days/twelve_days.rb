# Twelve Days
class TwelveDaysSong
  CARDINALS = {
    1 => 'first',
    2 => 'second',
    3 => 'third',
    4 => 'fourth',
    5 => 'fifth',
    6 => 'sixth',
    7 => 'seventh',
    8 => 'eighth',
    9 => 'ninth',
    10 => 'tenth',
    11 => 'eleventh',
    12 => 'twelfth'
  }.freeze

  MID_PARTIALS = {
    2 => 'two Turtle Doves',
    3 => 'three French Hens',
    4 => 'four Calling Birds',
    5 => 'five Gold Rings',
    6 => 'six Geese-a-Laying',
    7 => 'seven Swans-a-Swimming',
    8 => 'eight Maids-a-Milking',
    9 => 'nine Ladies Dancing',
    10 => 'ten Lords-a-Leaping',
    11 => 'eleven Pipers Piping',
    12 => 'twelve Drummers Drumming'
  }.freeze

  def verses(first = 1, last = 12)
    first.upto(last).map { |verse_num| verse(verse_num) }.join("\n") + "\n"
  end
  alias sing verses

  def verse(verse_num)
    if verse_num == 1
      "#{head(verse_num)}, #{tail(verse_num)}"
    else
      "#{head(verse_num)}, #{mid(verse_num)}, #{tail(verse_num)}"
    end
  end

  private

  def head(verse_num)
    "On the #{CARDINALS[verse_num]} day of Christmas my true love gave to me"
  end

  def mid(verse_num)
    verse_num.downto(2).map { |num| MID_PARTIALS[num] }.join(', ')
  end

  def tail(verse_num)
    "#{'and ' if verse_num > 1}a Partridge in a Pear Tree.\n"
  end
end
