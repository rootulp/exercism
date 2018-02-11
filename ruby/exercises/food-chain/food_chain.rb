# Food Chain Song
class FoodChainSong
  ANIMALS = {
    1 => :fly,
    2 => :spider,
    3 => :bird,
    4 => :cat,
    5 => :dog,
    6 => :goat,
    7 => :cow,
    8 => :horse
  }.freeze

  SECOND_LINES = {
    fly: '',
    spider: "It wriggled and jiggled and tickled inside her.\n",
    bird: "How absurd to swallow a bird!\n",
    cat: "Imagine that, to swallow a cat!\n",
    dog: "What a hog, to swallow a dog!\n",
    goat: "Just opened her throat and swallowed a goat!\n",
    cow: "I don't know how she swallowed a cow!\n",
    horse: "She's dead, of course!\n"
  }.freeze

  def verses(start = 1, stop = 8)
    (start..stop).map { |verse_num| verse(verse_num) }.join("\n") + "\n"
  end
  alias sing verses

  def verse(verse_num)
    animal = ANIMALS[verse_num]
    return head(animal) if animal == :horse
    head(animal) + middle_lines(verse_num) + tail
  end

  def head(animal)
    ["I know an old lady who swallowed a #{animal}.",
     SECOND_LINES[animal]].join("\n")
  end

  def middle_lines(verse_num)
    verse_num.downto(1).map do |curr|
      next if curr == 1
      animal1 = ANIMALS[curr]
      animal2 = ANIMALS[curr - 1]
      middle_line(animal1, animal2)
    end.join("\n")
  end

  def middle_line(animal1, animal2)
    "She swallowed the #{animal1} to catch the #{animal2}#{spider(animal2)}."
  end

  def spider(animal2)
    ' that wriggled and jiggled and tickled inside her' if animal2 == :spider
  end

  def tail
    "I don't know why she swallowed the fly. Perhaps she'll die.\n"
  end
end
