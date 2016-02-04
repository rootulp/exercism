# School
class School
  attr_accessor :db
  def initialize
    @db = Hash.new([])
  end

  def add(name, grade)
    db[grade] = [] unless db.key?(grade)
    (db[grade] << name).sort!
  end

  def grade(grade)
    db[grade]
  end

  def to_hash
    Hash[db.sort]
  end
end
