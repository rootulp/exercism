class BeerSong

  def verse(num)
    prefix(num) + suffix(num)
  end

  def verses(start, stop)
    result = ""
    start.downto(stop).each do |verse_num|
      result += verse(verse_num) + "\n"
    end
    result
  end

  def sing
    verses(99, 0)
  end

  def prefix(num)
    "#{num_or_No_more(num)} #{bottle_or_bottles(num)} of beer on the wall, "\
    "#{num_or_no_more(num)} #{bottle_or_bottles(num)} of beer.\n"
  end

  def suffix(num)
    if num == 0
      "Go to the store and buy some more, 99 bottles of beer on the wall.\n"
    else
      "Take #{one_or_it(num)} down and pass it around, #{num_or_no_more(num-1)} "\
      "#{bottle_or_bottles(num-1)} of beer on the wall.\n"
    end
  end

  def bottle_or_bottles(num)
    num == 1 ? "bottle" : "bottles"
  end

  def one_or_it(num)
    num != 1 ? "one" : "it"
  end

  def num_or_No_more(num)
    num != 0 ? num : "No more"
  end

  def num_or_no_more(num)
    num != 0 ? num : "no more"
  end
end