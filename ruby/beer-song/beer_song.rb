# Beer Song
class BeerSong
  def verses(start = 99, stop = 0)
    start.downto(stop).map do |verse_num|
      verse(verse_num)
    end.join("\n") + "\n"
  end
  alias sing verses

  def verse(num)
    [prefix(num), suffix(num)].join("\n") + "\n"
  end

  private

  def prefix(num)
    "#{quantity(num).capitalize} #{container(num)} of beer on the wall, "\
    "#{quantity(num)} #{container(num)} of beer."
  end

  def suffix(num)
    return last_line if num == 0
    "Take #{cardinality(num)} down and pass it around, #{quantity(num - 1)} "\
      "#{container(num - 1)} of beer on the wall."
  end

  def last_line
    'Go to the store and buy some more, 99 bottles of beer on the wall.'
  end

  def container(num)
    num == 1 ? 'bottle' : 'bottles'
  end

  def quantity(num)
    num == 0 ? 'no more' : num.to_s
  end

  def cardinality(num)
    num == 1 ? 'it' : 'one'
  end
end
